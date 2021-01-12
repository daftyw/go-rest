package rest

import "github.com/gin-gonic/gin"

// Start rest server
func StartServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	gin.SetMode(gin.ReleaseMode)
	r.Run()
}
