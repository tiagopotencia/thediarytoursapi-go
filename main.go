package main

import (
	"github.com/gin-gonic/gin"
	"git.heroku.com/thediarytoursapi-go/routes"
	"github.com/itsjamie/gin-cors"
	"time"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		//log.Fatal("$PORT must be set")
	}

	r := gin.New()
	r.Use(gin.Logger())

	// Apply the middleware to the router (works with groups too)
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	routes.SetRoutes(r)

	r.Run(":" + port)

}