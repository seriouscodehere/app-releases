package auth_login_db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateLoginTables(db *sql.DB) error {
	// Create login_otps table (for login verification)
	createLoginOTPsTable := `
	CREATE TABLE IF NOT EXISTS login_otps (
		email TEXT NOT NULL,
		code TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err := db.Exec(createLoginOTPsTable)
	if err != nil {
		return fmt.Errorf("failed to create login_otps table: %v", err)
	}

	// Login OTP requests table
	createLoginRequestTable := `
	CREATE TABLE IF NOT EXISTS login_otp_requests (
		email TEXT NOT NULL,
		request_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err = db.Exec(createLoginRequestTable)
	if err != nil {
		return fmt.Errorf("failed to create login_otp_requests table: %v", err)
	}

	// Login OTP cooldown table
	createLoginCooldownTable := `
	CREATE TABLE IF NOT EXISTS login_otp_cooldowns (
		email TEXT PRIMARY KEY,
		cooldown_until DATETIME,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err = db.Exec(createLoginCooldownTable)
	if err != nil {
		return fmt.Errorf("failed to create login_otp_cooldowns table: %v", err)
	}

	log.Println("Login tables created/verified")
	return nil
}
