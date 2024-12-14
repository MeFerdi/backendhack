package main

import (
	"log"

	"trial/controllers"
	"trial/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("Initializing database connection...")
	initializers.ConnectToDb()
	log.Println("Running database migrations...")
	initializers.SyncDb()
}

func main() {
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://safepulsefrontend.vercel.app"}, // Your frontend URL
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Define routes
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.POST("/satellite", controllers.Satelitte)
	router.POST("/satelitteLogin", controllers.SatelitteLogin)
	router.POST("/satelitteDashboard/add", controllers.DonateBlood)
	router.GET("/region", controllers.Region)
	router.GET("/donorPage/:userID", controllers.GetUserDonations)

	// Start the server
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
