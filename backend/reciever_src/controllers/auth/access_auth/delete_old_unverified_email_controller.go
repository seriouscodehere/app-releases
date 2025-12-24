package access_auth_controller

import (
	"database/sql"
	"log"
)

func AutoDeleteUnverifiedUsers(db *sql.DB) {
	_, err := db.Exec(`
		DELETE FROM users
		WHERE verified = 0 
		AND created_at <= datetime('now', '-24 hours')
	`)
	if err != nil {
		log.Println("Auto delete failed:", err)
	}
}
