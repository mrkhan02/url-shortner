package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mrkhan02/url-shortner-api/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	routes.ResRoutes(router)
	routes.ShortRoutes(router)
	router.Run()
}
