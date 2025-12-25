package user_profile_images_senders_controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

func getImageByType(c *gin.Context, imageType string) {
	uid := c.Query("uid")
	username := c.Query("username")

	if uid == "" || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uid and username required"})
		return
	}

	var url string

	err := DB.QueryRow(`
		SELECT url
		FROM user_profile_images
		WHERE uid = ? AND username = ? AND image_type = ?
		LIMIT 1
	`, uid, username, imageType).Scan(&url)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func GetProfilePhoto(c *gin.Context) {
	getImageByType(c, "profile")
}

func GetCoverImage(c *gin.Context) {
	getImageByType(c, "cover")
}
