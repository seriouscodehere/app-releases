package login_models

import (
	"database/sql"
	"errors"
	"time"
)

func GetUserIDByEmailOrUsername(db *sql.DB, login string) (int, error) {
	var id int
	err := db.QueryRow(`SELECT id FROM users WHERE email=? OR username=?`, login, login).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("user not found")
		}
		return 0, err
	}
	return id, nil
}

func GetStoredPasswordByEmail(db *sql.DB, email string) (string, error) {
	var password sql.NullString
	err := db.QueryRow("SELECT password FROM users WHERE email=?", email).Scan(&password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	if !password.Valid || password.String == "" {
		return "", errors.New("password not set")
	}
	return password.String, nil
}

func SaveLoginOTP(db *sql.DB, email, code string) error {
	_, _ = db.Exec("DELETE FROM login_otps WHERE email=?", email)
	_, err := db.Exec("INSERT INTO login_otps (email, code) VALUES (?, ?)", email, code)
	return err
}

func GetLoginOTP(db *sql.DB, email string) (string, time.Time, error) {
	var code string
	var created time.Time
	err := db.QueryRow("SELECT code, created_at FROM login_otps WHERE email=? ORDER BY created_at DESC LIMIT 1", email).Scan(&code, &created)
	return code, created, err
}

func DeleteLoginOTP(db *sql.DB, email string) error {
	_, err := db.Exec("DELETE FROM login_otps WHERE email=?", email)
	return err
}

func AddLoginOTPRequest(db *sql.DB, email string) error {
	_, err := db.Exec("INSERT INTO login_otp_requests (email) VALUES (?)", email)
	return err
}

func CountLoginRequestsLastHour(db *sql.DB, email string) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM login_otp_requests WHERE email=? AND request_time > datetime('now','-1 hour')", email).Scan(&count)
	return count, err
}

func SetLoginCooldown(db *sql.DB, email string, until time.Time) error {
	_, err := db.Exec("INSERT OR REPLACE INTO login_otp_cooldowns (email, cooldown_until) VALUES (?, ?)", email, until)
	return err
}

func GetLoginCooldown(db *sql.DB, email string) (time.Time, error) {
	var t time.Time
	err := db.QueryRow("SELECT cooldown_until FROM login_otp_cooldowns WHERE email=?", email).Scan(&t)
	return t, err
}
