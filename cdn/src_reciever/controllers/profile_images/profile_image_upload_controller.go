package profile_image_upload_controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"cdn/src_reciever/config"
	"cdn/src_reciever/mapping"

	"github.com/gin-gonic/gin"
)

var allowedFormats = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
}

func uploadImage(c *gin.Context, imageType string) {
	username := c.PostForm("username")
	uid := c.PostForm("uid")

	if username == "" || uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image not provided"})
		return
	}

	if file.Size > config.MaxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedFormats[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file format"})
		return
	}

	saveDir, err := mapping.EnsureImagePath(uid, username, imageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create directory"})
		return
	}

	finalFileName := imageType + ext
	finalPath := filepath.Join(saveDir, finalFileName)
	tempPath := finalPath + ".tmp"

	// save new image to temp file first
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		os.Remove(tempPath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
		return
	}

	// delete all existing images except temp
	files, err := os.ReadDir(saveDir)
	if err != nil {
		os.Remove(tempPath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read directory"})
		return
	}

	for _, f := range files {
		if f.Name() != filepath.Base(tempPath) {
			_ = os.Remove(filepath.Join(saveDir, f.Name()))
		}
	}

	// rename temp to final
	if err := os.Rename(tempPath, finalPath); err != nil {
		os.Remove(tempPath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to finalize image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
		"file":    finalFileName,
		"url":     "/media/" + uid + "/" + username + "/" + imageType + "/" + finalFileName,
	})
}

func UploadProfilePhoto(c *gin.Context) {
	uploadImage(c, "profile")
}

func UploadCoverImage(c *gin.Context) {
	uploadImage(c, "cover")
}
