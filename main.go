package main

import (
	"lunar/conf"
	"lunar/connectors"
	"lunar/routing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load the configuration
	cfg := conf.LoadConfig("conf/dbconfig.yml")

	// Initialize the database
	connectors.InitDatabase(&cfg.Database)

	port := "8080"

	// Initialize the router
	r := routing.RouterInstance(&gin.Context{})

	// Defines cors config
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST"}
	r.Use(cors.New(config))


	// Run the server
	r.Run(":" + port)
}
