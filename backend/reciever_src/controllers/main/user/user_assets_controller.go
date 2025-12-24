package user_assets_controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sraraa/db"
	user_models "sraraa/reciever_src/models/user"

	"github.com/gin-gonic/gin"
)

const CDN_UPLOAD_URL = "http://localhost:8090/api/upload/image"

type CDNResponse struct {
	File    string `json:"file"`
	Message string `json:"message"`
	URL     string `json:"url"`
}

// Upload Image
func UploadImage(c *gin.Context) {
	// Get the session token from cookie or header
	token := c.GetHeader("Authorization")
	if token == "" {
		token, _ = c.Cookie("session_token")
	}

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Validate session and get user claims
	claims, err := user_models.ValidateSessionToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
		return
	}

	// Get the image file from form
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image file is required"})
		return
	}
	defer file.Close()

	// Get the type from form (e.g., "profile", "cover", "post", etc.)
	imageType := c.PostForm("type")
	if imageType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type is required"})
		return
	}

	// Prepare multipart form to send to CDN
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("username", claims.Username)
	_ = writer.WriteField("uid", claims.UID)
	_ = writer.WriteField("type", imageType)

	part, err := writer.CreateFormFile("image", header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create form file"})
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to copy file"})
		return
	}

	err = writer.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to close writer"})
		return
	}

	// Send request to CDN API
	req, err := http.NewRequest("POST", CDN_UPLOAD_URL, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create CDN request"})
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload to CDN"})
		return
	}
	defer resp.Body.Close()

	// Parse CDN response
	var cdnResponse struct {
		File    string `json:"file"`
		Message string `json:"message"`
		URL     string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&cdnResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse CDN response"})
		return
	}

	if cdnResponse.URL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid CDN response"})
		return
	}

	// Prepend base URL to make full URL
	fullURL := fmt.Sprintf("http://localhost:8090%s", cdnResponse.URL)

	// Save full URL to database
	err = user_models.SaveUserImage(db.DB, claims.UID, claims.Username, imageType, fullURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"image_url": fullURL,
		"type":      imageType,
		"uid":       claims.UID,
		"username":  claims.Username,
	})
}

// GetImage retrieves image URL for a specific user and type
func GetImage(c *gin.Context) {
	// Accept either uid or username in query
	uid := c.Query("uid")
	username := c.Query("username")
	imageType := c.Param("type")

	if uid == "" && username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uid or username is required"})
		return
	}

	if imageType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type is required"})
		return
	}

	// Get image URL from database
	imageData, err := user_models.GetUserImage(db.DB, uid, username, imageType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"image_url": imageData["image_url"],
		"type":      imageData["type"],
		"uid":       imageData["uid"],
		"username":  imageData["username"],
	})
}

// GetAllUserImages retrieves all images for a specific user
func GetAllUserImages(c *gin.Context) {
	// Accept either uid or username in query
	uid := c.Query("uid")
	username := c.Query("username")

	if uid == "" && username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uid or username is required"})
		return
	}

	// Get all images from database
	images, err := user_models.GetAllUserImages(db.DB, uid, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve images"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"uid":      uid,
		"username": username,
		"images":   images,
	})
}

// DeleteImage removes an image record from database
func DeleteImage(c *gin.Context) {
	// Get the session token from cookie or header
	token := c.GetHeader("Authorization")
	if token == "" {
		token, _ = c.Cookie("session_token")
	}

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Validate session and get user claims
	claims, err := user_models.ValidateSessionToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
		return
	}

	imageType := c.Param("type")
	if imageType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type is required"})
		return
	}

	// Delete image from database
	err = user_models.DeleteUserImage(db.DB, claims.UID, imageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("Image of type '%s' deleted successfully", imageType),
	})
}
