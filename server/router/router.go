package Router

import (
	"github.com/gin-gonic/gin"
	. "grs/controller"
	. "grs/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(CorsMiddleware())
	r.POST("/api/auth/register", Register)
	r.POST("/api/auth/login", Login)
	r.GET("/api/auth/info", AuthMiddleware(), Info)
	r.GET("/api/auth/UpdateProfile", AuthMiddleware(), UpdateProfile)
	return r
}
