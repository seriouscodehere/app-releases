package session_models

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("2669b02cb3dfd75154206495c9d460e09080cc398a5c02abe2ede88e883dfd26")

type SessionClaims struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Verified bool   `json:"verified"`
	UID      string `json:"uid"`
	jwt.RegisteredClaims
}

func CreateSession(db *sql.DB, userID int, duration time.Duration, userAgent, ip string) (string, error) {
	if userID <= 0 {
		return "", errors.New("invalid user ID")
	}

	var email, username, fullname, uid string
	var verified bool
	err := db.QueryRow(`SELECT email, username, fullname, verified, uid FROM users WHERE id=?`, userID).
		Scan(&email, &username, &fullname, &verified, &uid)
	if err != nil {
		return "", err
	}

	if uid == "" {
		return "", errors.New("user does not have a UID assigned")
	}

	nonce := make([]byte, 16)
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	claims := SessionClaims{
		UserID:   userID,
		Email:    email,
		Username: username,
		Fullname: fullname,
		Verified: verified,
		UID:      uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        hex.EncodeToString(nonce),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	_, err = db.Exec(`
		INSERT INTO sessions (uid, session_token, user_agent, ip_address, expires_at)
		VALUES (?, ?, ?, ?, ?)`,
		uid, signedToken, userAgent, ip, time.Now().Add(duration))
	if err != nil {
		log.Println("CreateSession insert failed:", err)
		return "", err
	}

	return signedToken, nil
}

func DeleteSession(db *sql.DB, token string) error {
	res, err := db.Exec(`DELETE FROM sessions WHERE session_token=?`, token)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	log.Printf("Deleted %d session(s) with token=%s\n", rows, token)
	return nil
}

func DeleteAllSessions(db *sql.DB, userID int) error {
	var uid string
	err := db.QueryRow(`SELECT uid FROM users WHERE id=?`, userID).Scan(&uid)
	if err != nil {
		return err
	}

	if uid == "" {
		return errors.New("user does not have a UID")
	}

	res, err := db.Exec(`DELETE FROM sessions WHERE uid=?`, uid)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	log.Printf("Deleted %d session(s) for UID=%s\n", rows, uid)
	return nil
}

func DeleteAllSessionsByUID(db *sql.DB, uid string) error {
	if uid == "" {
		return errors.New("uid cannot be empty")
	}

	res, err := db.Exec(`DELETE FROM sessions WHERE uid=?`, uid)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	log.Printf("Deleted %d session(s) for UID=%s\n", rows, uid)
	return nil
}

func GetSessionsByUID(db *sql.DB, uid string) ([]map[string]interface{}, error) {
	if uid == "" {
		return nil, errors.New("uid cannot be empty")
	}

	rows, err := db.Query(`
		SELECT session_token, user_agent, ip_address, created_at, expires_at 
		FROM sessions 
		WHERE uid=? AND expires_at > datetime('now')
		ORDER BY created_at DESC
	`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []map[string]interface{}
	for rows.Next() {
		var token, userAgent, ipAddress string
		var createdAt, expiresAt time.Time
		if err := rows.Scan(&token, &userAgent, &ipAddress, &createdAt, &expiresAt); err != nil {
			continue
		}
		sessions = append(sessions, map[string]interface{}{
			"session_token": token,
			"user_agent":    userAgent,
			"ip_address":    ipAddress,
			"created_at":    createdAt,
			"expires_at":    expiresAt,
		})
	}
	return sessions, nil
}

func CountActiveSessions(db *sql.DB, uid string) (int, error) {
	if uid == "" {
		return 0, errors.New("uid cannot be empty")
	}

	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) 
		FROM sessions 
		WHERE uid=? AND expires_at > datetime('now')
	`, uid).Scan(&count)

	return count, err
}

func ValidateSessionToken(tokenStr string) (*SessionClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &SessionClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*SessionClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid session token")
}
