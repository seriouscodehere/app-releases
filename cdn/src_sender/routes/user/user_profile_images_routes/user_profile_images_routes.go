package user_profile_images_routes

import (
	"cdn/src_sender/controllers/users/user_profile_images_senders_controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserProfileImageRoutes(router *gin.Engine) {
	router.GET("/api/profile-photo", user_profile_images_senders_controller.GetProfilePhoto)
	router.GET("/api/cover", user_profile_images_senders_controller.GetCoverImage)
}
