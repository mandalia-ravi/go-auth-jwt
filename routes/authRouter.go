package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", handler.Signup())
	incomingRoutes.POST("/users/login", handler.Login())
}
