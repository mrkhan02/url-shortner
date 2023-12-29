package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrkhan02/url-shortner-api/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.ResRoutes(router)
	routes.ShortRoutes(router)
	router.Run()
}
