package auth_signup_db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateSignupTables(db *sql.DB) error {
	// Create otps table (for registration)
	createOTPsTable := `
	CREATE TABLE IF NOT EXISTS signup_otps (
		email TEXT NOT NULL,
		code TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err := db.Exec(createOTPsTable)
	if err != nil {
		return fmt.Errorf("failed to create otps table: %v", err)
	}

	// OTP requests table (for registration)
	createRequestTable := `
	CREATE TABLE IF NOT EXISTS signup_otp_requests (
		email TEXT NOT NULL,
		request_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err = db.Exec(createRequestTable)
	if err != nil {
		return fmt.Errorf("failed to create otp_requests table: %v", err)
	}

	// OTP cooldown table (for registration)
	createCooldownTable := `
	CREATE TABLE IF NOT EXISTS signup_otp_cooldowns (
		email TEXT PRIMARY KEY,
		cooldown_until DATETIME,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err = db.Exec(createCooldownTable)
	if err != nil {
		return fmt.Errorf("failed to create signup_otp_cooldowns table: %v", err)
	}

	log.Println("Signup tables created/verified")
	return nil
}
