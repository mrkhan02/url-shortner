package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrkhan02/url-shortner-api/routes"
)

func main() {
	router := gin.Default()

	// Enable CORS for all origins
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	})

	routes.ResRoutes(router)
	routes.ShortRoutes(router)

	router.Run()
}
