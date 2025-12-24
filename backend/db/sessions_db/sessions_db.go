package sessions_db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateSessionsTable(db *sql.DB) error {
	createSessionsTable := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid TEXT NOT NULL,
		session_token TEXT UNIQUE NOT NULL,
		user_agent TEXT,
		ip_address TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		expires_at DATETIME,
		FOREIGN KEY(uid) REFERENCES users(uid) ON DELETE CASCADE
	);
	`

	_, err := db.Exec(createSessionsTable)
	if err != nil {
		return fmt.Errorf("failed to create sessions table: %v", err)
	}

	log.Println("Sessions table created/verified")
	return nil
}
