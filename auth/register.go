package auth

import (
	"github.com/gin-gonic/gin"
	"go-learning-demo/auth/controller"
)

func RegisterHttpEndpoints(router *gin.Engine, controller *controller.AuthController) {

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", controller.SignUp)
		authEndpoints.POST("/sign-in", controller.SignIn)
	}
}
