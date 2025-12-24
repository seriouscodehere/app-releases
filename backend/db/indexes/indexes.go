package indexes

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateAllIndexes(db *sql.DB) error {
	// Index for users table
	userIndexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);`,
		`CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);`,
		`CREATE INDEX IF NOT EXISTS idx_users_uid ON users(uid);`,
	}

	// Index for sessions table
	sessionIndexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_sessions_uid ON sessions(uid);`,
		`CREATE INDEX IF NOT EXISTS idx_sessions_token ON sessions(session_token);`,
		`CREATE INDEX IF NOT EXISTS idx_sessions_expires ON sessions(expires_at);`,
	}

	// Index for otp tables
	otpIndexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_otps_email ON signup_otps(email);`,
		`CREATE INDEX IF NOT EXISTS idx_login_otps_email ON login_otps(email);`,
		`CREATE INDEX IF NOT EXISTS idx_password_otps_email ON password_reset_otps(email);`,
	}

	// Index for request tables
	requestIndexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_signup_otp_requests_email_time ON signup_otp_requests(email, request_time);`,
		`CREATE INDEX IF NOT EXISTS idx_login_requests_email_time ON login_otp_requests(email, request_time);`,
		`CREATE INDEX IF NOT EXISTS idx_password_requests_email_time ON password_reset_requests(email, request_time);`,
	}

	// Index for user_images table
	imageIndexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_user_images_uid ON user_images(uid);`,
		`CREATE INDEX IF NOT EXISTS idx_user_images_type ON user_images(type);`,
		`CREATE INDEX IF NOT EXISTS idx_user_images_username ON user_images(username);`,
	}

	// Execute all indexes
	allIndexes := [][]string{
		userIndexes,
		sessionIndexes,
		otpIndexes,
		requestIndexes,
		imageIndexes,
	}

	for _, indexGroup := range allIndexes {
		for _, indexSQL := range indexGroup {
			_, err := db.Exec(indexSQL)
			if err != nil {
				return fmt.Errorf("failed to create index: %v\nSQL: %s", err, indexSQL)
			}
		}
	}

	log.Println("All indexes created/verified")
	return nil
}
