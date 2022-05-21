package Router

import (
	. "github.com/Ran-Xing/ATS/server/controller"
	. "github.com/Ran-Xing/ATS/server/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(CorsMiddleware())
	r.POST("/api/auth/register", Register)
	r.POST("/api/auth/login", Login)
	r.GET("/api/auth/info", AuthMiddleware(), Info)
	r.GET("/api/auth/UpdateProfile", AuthMiddleware(), UpdateProfile)
	return r
}
