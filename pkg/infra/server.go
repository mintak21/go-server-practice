package infra

import (
	"github.com/gin-gonic/gin"
)

// RunServer is now Quick Starting Code with gin-gonic
func RunServer() error {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	return nil
}
