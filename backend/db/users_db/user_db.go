package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateUsersTable(db *sql.DB) error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid TEXT UNIQUE,
		email TEXT UNIQUE NOT NULL,
		username TEXT UNIQUE,
		password TEXT,
		fullname TEXT,
		verified BOOLEAN DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(createUserTable)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	log.Println("Users table created/verified")
	return nil
}
