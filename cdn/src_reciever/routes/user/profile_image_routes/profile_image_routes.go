package profile_image_routes

import (
	profile_image_upload_controller "cdn/src_reciever/controllers/profile_images"

	"github.com/gin-gonic/gin"
)

func UploadRoutes(router *gin.Engine) {
	router.POST("/api/upload/image/profile-photo-image", profile_image_upload_controller.UploadProfilePhoto)
	router.POST("/api/upload/image/profile-cover-image", profile_image_upload_controller.UploadCoverImage)
}
