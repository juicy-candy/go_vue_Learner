package router

import (
	"ginvue/pkg/controller"
	"ginvue/pkg/midware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("api/auth/info", midware.AuthMiddleware(), controller.Info)
	return r
}
