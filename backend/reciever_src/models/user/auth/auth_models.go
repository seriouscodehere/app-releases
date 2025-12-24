package auth_models

import (
	"database/sql"
	"errors"
	"time"

	auth_utils "sraraa/reciever_src/utils/auth"
)

func SetPassword(db *sql.DB, email, password string) error {
	if err := auth_utils.ValidatePassword(password); err != nil {
		return err
	}
	_, err := db.Exec("UPDATE users SET password=? WHERE email=?", password, email)
	return err
}

func SetUsername(db *sql.DB, email, username string) error {
	if err := auth_utils.ValidateUsername(username); err != nil {
		return err
	}
	_, err := db.Exec("UPDATE users SET username=? WHERE email=?", username, email)
	return err
}

func SetFullname(db *sql.DB, email, fullname string) error {
	if err := auth_utils.ValidateFullname(fullname); err != nil {
		return err
	}
	_, err := db.Exec("UPDATE users SET fullname=? WHERE email=?", fullname, email)
	return err
}

func UsernameExists(db *sql.DB, username string) (bool, error) {
	if username == "" {
		return false, nil
	}

	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM users WHERE username=?`, username).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func EmailExists(db *sql.DB, email string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email=?", email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func HasAllRequiredFields(db *sql.DB, email string) (bool, time.Time, error) {
	var username, password, fullname sql.NullString
	var updated time.Time

	err := db.QueryRow(`
		SELECT username, password, fullname, created_at
		FROM users WHERE email=?
	`, email).Scan(&username, &password, &fullname, &updated)

	if err != nil {
		return false, time.Time{}, err
	}

	if !username.Valid || !password.Valid || !fullname.Valid {
		return false, updated, nil
	}

	return true, updated, nil
}

func HasAllRequiredFieldsForLogin(db *sql.DB, email string) (bool, error) {
	var username, password, fullname sql.NullString

	err := db.QueryRow(`
		SELECT username, password, fullname
		FROM users
		WHERE email=?
	`, email).Scan(&username, &password, &fullname)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	if username.Valid && password.Valid && fullname.Valid {
		return true, nil
	}

	return false, nil
}
