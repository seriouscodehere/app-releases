// package models
package models

import "time"

type UserProfileImage struct {
	ID        int64     `db:"id"`
	UID       string    `db:"uid"`
	Username  string    `db:"username"`
	ImageType string    `db:"image_type"`
	FileName  string    `db:"file_name"`
	URL       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
