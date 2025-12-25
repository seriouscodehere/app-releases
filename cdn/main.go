package main

import (
	"log"

	"cdn/cors"
	db "cdn/db/main"
	profile_image_upload_controller "cdn/src_reciever/controllers/profile_images"
	"cdn/src_reciever/routes/user/profile_image_routes"
	"cdn/src_reciever/static"
	"cdn/src_sender/controllers/users/user_profile_images_senders_controller"
	"cdn/src_sender/routes/user/user_profile_images_routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn, err := db.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	profile_image_upload_controller.SetDB(dbConn)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(cors.AllowLocalHTML())

	static.RegisterStaticRoutes(r)
	profile_image_routes.UploadRoutes(r)
	user_profile_images_senders_controller.SetDB(dbConn)
	user_profile_images_routes.RegisterUserProfileImageRoutes(r)

	r.Run(":8090")
}
