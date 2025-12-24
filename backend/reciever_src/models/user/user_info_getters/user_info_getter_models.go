package user_info_getter_models

import (
	"database/sql"
	"errors"
	"sraraa/db"
	"time"
)

func GetUserEmailByUID(uid string) (string, error) {
	if uid == "" {
		return "", errors.New("uid cannot be empty")
	}
	var value sql.NullString
	err := db.DB.QueryRow(`SELECT email FROM users WHERE uid=?`, uid).Scan(&value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	if value.Valid {
		return value.String, nil
	}
	return "", nil
}

func GetUsernameByUID(uid string) (string, error) {
	if uid == "" {
		return "", errors.New("uid cannot be empty")
	}
	var value sql.NullString
	err := db.DB.QueryRow(`SELECT username FROM users WHERE uid=?`, uid).Scan(&value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	if value.Valid {
		return value.String, nil
	}
	return "", nil
}

func GetFullnameByUID(uid string) (string, error) {
	if uid == "" {
		return "", errors.New("uid cannot be empty")
	}
	var value sql.NullString
	err := db.DB.QueryRow(`SELECT fullname FROM users WHERE uid=?`, uid).Scan(&value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	if value.Valid {
		return value.String, nil
	}
	return "", nil
}

func GetPasswordByUID(uid string) (string, error) {
	if uid == "" {
		return "", errors.New("uid cannot be empty")
	}
	var value sql.NullString
	err := db.DB.QueryRow(`SELECT password FROM users WHERE uid=?`, uid).Scan(&value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	if value.Valid {
		return value.String, nil
	}
	return "", nil
}

func GetUserVerifiedByUID(uid string) (bool, error) {
	if uid == "" {
		return false, errors.New("uid cannot be empty")
	}
	var verified bool
	err := db.DB.QueryRow(`SELECT verified FROM users WHERE uid=?`, uid).Scan(&verified)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("user not found")
		}
		return false, err
	}
	return verified, nil
}

func GetUserIDByUID(uid string) (int, error) {
	if uid == "" {
		return 0, errors.New("uid cannot be empty")
	}
	var id int
	err := db.DB.QueryRow(`SELECT id FROM users WHERE uid=?`, uid).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("user not found")
		}
		return 0, err
	}
	return id, nil
}

// Password Reset Models
func SavePasswordResetOTP(db *sql.DB, email, code string) error {
	_, _ = db.Exec("DELETE FROM password_reset_otps WHERE email=?", email)
	_, err := db.Exec("INSERT INTO password_reset_otps (email, code) VALUES (?, ?)", email, code)
	return err
}

func GetPasswordResetOTP(db *sql.DB, email string) (string, time.Time, error) {
	var code string
	var created time.Time
	err := db.QueryRow(
		"SELECT code, created_at FROM password_reset_otps WHERE email=? ORDER BY created_at DESC LIMIT 1",
		email,
	).Scan(&code, &created)
	return code, created, err
}

func DeletePasswordResetOTP(db *sql.DB, email string) error {
	_, err := db.Exec("DELETE FROM password_reset_otps WHERE email=?", email)
	return err
}

func AddPasswordResetRequest(db *sql.DB, email string) error {
	_, err := db.Exec("INSERT INTO password_reset_requests (email) VALUES (?)", email)
	return err
}

func CountPasswordResetRequestsLastHour(db *sql.DB, email string) (int, error) {
	var count int
	err := db.QueryRow(
		"SELECT COUNT(*) FROM password_reset_requests WHERE email=? AND request_time > datetime('now','-1 hour')",
		email,
	).Scan(&count)
	return count, err
}

func SetPasswordResetCooldown(db *sql.DB, email string, until time.Time) error {
	_, err := db.Exec(
		"INSERT OR REPLACE INTO password_reset_cooldowns (email, cooldown_until) VALUES (?, ?)",
		email,
		until,
	)
	return err
}

func GetPasswordResetCooldown(db *sql.DB, email string) (time.Time, error) {
	var t time.Time
	err := db.QueryRow(
		"SELECT cooldown_until FROM password_reset_cooldowns WHERE email=?",
		email,
	).Scan(&t)
	return t, err
}
