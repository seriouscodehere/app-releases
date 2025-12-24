package user_images_models

import (
	"database/sql"
	"errors"
	"time"
)

func SaveUserImage(db *sql.DB, uid, username, imageType, imageURL string) error {
	if uid == "" || imageType == "" || imageURL == "" {
		return errors.New("uid, imageType, and imageURL are required")
	}

	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM user_images WHERE uid=? AND type=?`, uid, imageType).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err = db.Exec(`UPDATE user_images SET username=?, image_url=?, updated_at=CURRENT_TIMESTAMP WHERE uid=? AND type=?`,
			username, imageURL, uid, imageType)
	} else {
		_, err = db.Exec(`INSERT INTO user_images (uid, username, type, image_url) VALUES (?, ?, ?, ?)`,
			uid, username, imageType, imageURL)
	}

	return err
}

func GetUserImage(db *sql.DB, uid, username, imageType string) (map[string]interface{}, error) {
	if (uid == "" && username == "") || imageType == "" {
		return nil, errors.New("uid or username, and imageType are required")
	}

	var query string
	var args []interface{}

	if uid != "" {
		query = `SELECT uid, username, type, image_url FROM user_images WHERE uid=? AND type=?`
		args = []interface{}{uid, imageType}
	} else {
		query = `SELECT uid, username, type, image_url FROM user_images WHERE username=? AND type=?`
		args = []interface{}{username, imageType}
	}

	var resultUID, resultUsername, resultType, imageURL string
	err := db.QueryRow(query, args...).Scan(&resultUID, &resultUsername, &resultType, &imageURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("image not found")
		}
		return nil, err
	}

	return map[string]interface{}{
		"uid":       resultUID,
		"username":  resultUsername,
		"type":      resultType,
		"image_url": imageURL,
	}, nil
}

func GetAllUserImages(db *sql.DB, uid, username string) ([]map[string]interface{}, error) {
	if uid == "" && username == "" {
		return nil, errors.New("uid or username is required")
	}

	var query string
	var args []interface{}

	if uid != "" {
		query = `
			SELECT uid, username, type, image_url, created_at, updated_at 
			FROM user_images 
			WHERE uid=?
			ORDER BY created_at DESC
		`
		args = []interface{}{uid}
	} else {
		query = `
			SELECT uid, username, type, image_url, created_at, updated_at 
			FROM user_images 
			WHERE username=?
			ORDER BY created_at DESC
		`
		args = []interface{}{username}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []map[string]interface{}
	for rows.Next() {
		var resultUID, resultUsername, imageType, imageURL string
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&resultUID, &resultUsername, &imageType, &imageURL, &createdAt, &updatedAt); err != nil {
			continue
		}
		images = append(images, map[string]interface{}{
			"uid":        resultUID,
			"username":   resultUsername,
			"type":       imageType,
			"image_url":  imageURL,
			"created_at": createdAt,
			"updated_at": updatedAt,
		})
	}

	return images, nil
}

func DeleteUserImage(db *sql.DB, uid, imageType string) error {
	if uid == "" || imageType == "" {
		return errors.New("uid and imageType are required")
	}

	result, err := db.Exec(`DELETE FROM user_images WHERE uid=? AND type=?`, uid, imageType)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("image not found")
	}

	return nil
}
