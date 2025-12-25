package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB   *sql.DB
	once sync.Once
)

// InitDB initializes sqlite with WAL mode and shm
func InitDB() *sql.DB {
	once.Do(func() {
		dbPath := "cdn.db"

		var err error
		DB, err = sql.Open(
			"sqlite3",
			dbPath+"?_journal_mode=WAL&_foreign_keys=ON&_busy_timeout=5000",
		)
		if err != nil {
			log.Fatal("failed to open database:", err)
		}

		if err := DB.Ping(); err != nil {
			log.Fatal("failed to ping database:", err)
		}
	})
	return DB
}

// InitializeDatabase creates tables and indexes
func InitializeDatabase() (*sql.DB, error) {
	db := InitDB()

	if err := InitUserProfileImageTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *sql.DB {
	if DB == nil {
		return InitDB()
	}
	return DB
}

func CloseDB() {
	if DB != nil {
		_ = DB.Close()
	}
}

// Table + indexes in same place
func InitUserProfileImageTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS user_profile_images (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid TEXT NOT NULL,
		username TEXT NOT NULL,
		image_type TEXT NOT NULL,
		file_name TEXT NOT NULL,
		url TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(uid, image_type)
	);

	CREATE INDEX IF NOT EXISTS idx_user_profile_images_uid
	ON user_profile_images(uid);

	CREATE INDEX IF NOT EXISTS idx_user_profile_images_username
	ON user_profile_images(username);
	`
	_, err := db.Exec(query)
	return err
}
