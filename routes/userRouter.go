package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iravimandalia/go-auth-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", handler.GetUsers())
	incomingRoutes.GET("/users/:user_id", handler.GetUser)
}
