package go_auth_jwt

import (
	"github.com/gin-gonic/gin"
	routes "github.com/iravimandalia/go-auth-jwt/routes"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Success": "Access granted for api-2"})
	})

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
