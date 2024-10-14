package main

import (
	"Mereb-V1/config"
	"Mereb-V1/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	Port := os.Getenv("PORT")

	if Port == "" {
		Port = "5000"
	}

	router := gin.New()

	//  Manually add Logger and Recovery

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(config.CorsConfig)

	routes.PersonRoutes(router)

	router.Run(":" + Port)

}
