package user_image_db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateImagesTables(db *sql.DB) error {
	// Create user_images table for storing image URLs
	createUserImagesTable := `
	CREATE TABLE IF NOT EXISTS user_images (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid TEXT NOT NULL,
		username TEXT NOT NULL,
		type TEXT NOT NULL,
		image_url TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(uid, type),
		FOREIGN KEY(uid) REFERENCES users(uid) ON DELETE CASCADE
	);
	`

	_, err := db.Exec(createUserImagesTable)
	if err != nil {
		return fmt.Errorf("failed to create user_images table: %v", err)
	}

	log.Println("Images tables created/verified")
	return nil
}
