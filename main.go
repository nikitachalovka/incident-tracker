package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/student/incident-tracker/config"
	"github.com/student/incident-tracker/handlers"
	"github.com/student/incident-tracker/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/student/incident-tracker/docs"
)

// @title Incident Tracker API
// @version 1.0
// @description REST API for tracking and managing incidents.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter: Bearer followed by your JWT token.
func main() {
	err := godotenv.Load()
	log.Printf("DB_HOST=%s", os.Getenv("DB_HOST"))
log.Printf("DB_USER=%s", os.Getenv("DB_USER"))
log.Printf("DB_NAME=%s", os.Getenv("DB_NAME"))
log.Printf("DB_PASSWORD length=%d", len(os.Getenv("DB_PASSWORD")))
	if err != nil {
		log.Println("Warning: .env file not found; using environment variables")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not configured")
	}

	db := config.ConnectDatabase()

	authHandler := handlers.NewAuthHandler(db)
	incidentHandler := handlers.NewIncidentHandler(db)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// All incident routes require a valid JWT.
		incidents := api.Group("/incidents")
		incidents.Use(middleware.AuthMiddleware(jwtSecret))
		{
			incidents.POST("", incidentHandler.CreateIncident)
			incidents.GET("", incidentHandler.ListIncidents)
			incidents.GET("/:id", incidentHandler.GetIncident)
			incidents.PATCH("/:id/status", incidentHandler.UpdateIncidentStatus)
			incidents.DELETE("/:id", incidentHandler.DeleteIncident)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Incident Tracker API is running on port %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
