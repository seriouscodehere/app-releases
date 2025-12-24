package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"

	"sraraa/db/auth_login_db"
	"sraraa/db/auth_password_db"
	"sraraa/db/auth_signup_db"
	"sraraa/db/indexes"
	"sraraa/db/sessions_db"
	"sraraa/db/user_image_db"
	"sraraa/db/users_db"
)

var (
	DB   *sql.DB
	once sync.Once
)

// InitDB initializes the database connection (singleton pattern)
func InitDB() *sql.DB {
	once.Do(func() {
		dbPath := "users.db"
		var err error
		DB, err = sql.Open("sqlite3", dbPath+"?_journal_mode=WAL&_foreign_keys=ON")
		if err != nil {
			log.Fatal("Failed to open database:", err)
			return
		}

		// Test the connection
		if err := DB.Ping(); err != nil {
			log.Fatal("Failed to ping database:", err)
		}
	})
	return DB
}

// InitializeDatabase creates all tables and indexes
func InitializeDatabase() (*sql.DB, error) {
	// Initialize connection
	db := InitDB()

	// Create all tables in order of dependencies
	tables := []struct {
		name string
		fn   func(*sql.DB) error
	}{
		{"users", users_db.CreateUsersTable},
		{"sessions", sessions_db.CreateSessionsTable},
		{"signup auth", auth_signup_db.CreateSignupTables},
		{"login auth", auth_login_db.CreateLoginTables},
		{"password auth", auth_password_db.CreatePasswordTables},
		{"images", user_image_db.CreateImagesTables},
	}

	log.Println("Starting database initialization...")

	for _, table := range tables {
		log.Printf("Creating %s tables...", table.name)
		if err := table.fn(db); err != nil {
			return nil, fmt.Errorf("failed to create %s tables: %v", table.name, err)
		}
	}

	// Create indexes
	log.Println("Creating indexes...")
	if err := indexes.CreateAllIndexes(db); err != nil {
		return nil, fmt.Errorf("failed to create indexes: %v", err)
	}

	log.Println("Database initialized successfully")
	return db, nil
}

// GetDB returns the global database instance
func GetDB() *sql.DB {
	if DB == nil {
		return InitDB()
	}
	return DB
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// BeginTransaction starts a new database transaction
func BeginTransaction() (*sql.Tx, error) {
	return GetDB().Begin()
}

// ExecuteWithTransaction executes a function within a transaction
func ExecuteWithTransaction(fn func(*sql.Tx) error) error {
	tx, err := BeginTransaction()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
