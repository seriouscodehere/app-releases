package signup_models

import (
	"database/sql"
	"time"
)

func IsVerified(db *sql.DB, email string) (bool, error) {
	var verified bool
	err := db.QueryRow("SELECT verified FROM users WHERE email=?", email).Scan(&verified)
	return verified, err
}

func MarkVerified(db *sql.DB, email string) error {
	_, err := db.Exec("UPDATE users SET verified=1 WHERE email=?", email)
	return err
}

func SaveOTP(db *sql.DB, email, code string) error {
	_, _ = db.Exec("DELETE FROM signup_otps WHERE email=?", email)
	_, err := db.Exec("INSERT INTO signup_otps (email, code) VALUES (?, ?)", email, code)
	return err
}

func GetOTP(db *sql.DB, email string) (string, time.Time, error) {
	var code string
	var created time.Time
	err := db.QueryRow("SELECT code, created_at FROM signup_otps WHERE email=? ORDER BY created_at DESC LIMIT 1", email).Scan(&code, &created)
	return code, created, err
}

func DeleteOTP(db *sql.DB, email string) error {
	_, err := db.Exec("DELETE FROM signup_otps WHERE email=?", email)
	return err
}

func AddOTPRequest(db *sql.DB, email string) error {
	_, err := db.Exec("INSERT INTO signup_otp_requests (email) VALUES (?)", email)
	return err
}

func CountRequestsLastHour(db *sql.DB, email string) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM signup_otp_requests WHERE email=? AND request_time > datetime('now','-1 hour')", email).Scan(&count)
	return count, err
}

func SetCooldown(db *sql.DB, email string, until time.Time) error {
	_, err := db.Exec("INSERT OR REPLACE INTO signup_otp_cooldowns (email, cooldown_until) VALUES (?, ?)", email, until)
	return err
}

func GetCooldown(db *sql.DB, email string) (time.Time, error) {
	var t time.Time
	err := db.QueryRow("SELECT cooldown_until FROM signup_otp_cooldowns WHERE email=?", email).Scan(&t)
	return t, err
}
