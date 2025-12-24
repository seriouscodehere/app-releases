package onboard_models

import (
	"database/sql"
)

func CreateUser(db *sql.DB, email string) error {
	_, err := db.Exec("INSERT OR IGNORE INTO users (email) VALUES (?)", email)
	return err
}

func SetUniqueID(db *sql.DB, email, uid string) error {
	_, err := db.Exec(`UPDATE users SET uid=? WHERE email=?`, uid, email)
	return err
}

func HasUID(db *sql.DB, email string) (bool, error) {
	var uid sql.NullString
	err := db.QueryRow(`SELECT uid FROM users WHERE email=?`, email).Scan(&uid)
	if err != nil {
		return false, err
	}
	return uid.Valid && uid.String != "", nil
}

func UniqueIDExists(db *sql.DB, uid string) (bool, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM users WHERE uid=?`, uid).Scan(&count)
	return count > 0, err
}
