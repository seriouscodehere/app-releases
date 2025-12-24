package auth_password_db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreatePasswordTables(db *sql.DB) error {
	// Password reset OTPs
	createPasswordResetOTPs := `
	CREATE TABLE IF NOT EXISTS password_reset_otps (
		email TEXT NOT NULL,
		code TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err := db.Exec(createPasswordResetOTPs)
	if err != nil {
		return fmt.Errorf("failed to create password_reset_otps table: %v", err)
	}

	// Password reset requests
	createPasswordResetRequests := `
	CREATE TABLE IF NOT EXISTS password_reset_requests (
		email TEXT NOT NULL,
		request_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err = db.Exec(createPasswordResetRequests)
	if err != nil {
		return fmt.Errorf("failed to create password_reset_requests table: %v", err)
	}

	// Password reset cooldown
	createPasswordResetCooldown := `
	CREATE TABLE IF NOT EXISTS password_reset_cooldowns (
		email TEXT PRIMARY KEY,
		cooldown_until DATETIME,
		FOREIGN KEY(email) REFERENCES users(email) ON DELETE CASCADE
	);
	`

	_, err = db.Exec(createPasswordResetCooldown)
	if err != nil {
		return fmt.Errorf("failed to create password_reset_cooldowns table: %v", err)
	}

	log.Println("Password reset tables created/verified")
	return nil
}
