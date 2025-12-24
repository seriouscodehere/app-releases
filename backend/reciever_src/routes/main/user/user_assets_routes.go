package user_assets_routes

import (
	user_assets_controller "sraraa/reciever_src/controllers/main/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserAssetsRoutes(router *gin.Engine) {
	// Image upload and management routes
	imageRoutes := router.Group("/image")
	{
		imageRoutes.POST("", user_assets_controller.UploadImage)

		imageRoutes.GET("/all", user_assets_controller.GetAllUserImages)

		imageRoutes.GET("/:type", user_assets_controller.GetImage)

		imageRoutes.DELETE("/:type", user_assets_controller.DeleteImage)
	}

}
