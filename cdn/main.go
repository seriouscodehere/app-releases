package main

import (
	"cdn/src_reciever/routes/user/profile_image_routes"
	"cdn/src_reciever/static"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	static.RegisterStaticRoutes(r)

	profile_image_routes.UploadRoutes(r)

	r.Run(":8090")
}
