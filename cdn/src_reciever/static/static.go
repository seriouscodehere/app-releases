package static

import "github.com/gin-gonic/gin"

func RegisterStaticRoutes(r *gin.Engine) {
	r.Static("/media", "src_reciever/storage")
}
