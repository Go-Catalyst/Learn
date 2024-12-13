package main

import (
	"log"
	"os"

	"JWT/handlers"
	"JWT/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Println(".env file loaded successfully")
	}
}

func main() {
	r := gin.Default()

	// Public routes
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a protected route"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
