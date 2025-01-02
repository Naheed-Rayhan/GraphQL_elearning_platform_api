package main

import (
	"github.com/Naheed-Rayhan/graphql-api/config"
	"github.com/Naheed-Rayhan/graphql-api/infrastructure"
	"github.com/Naheed-Rayhan/graphql-api/infrastructure/database"
	"github.com/Naheed-Rayhan/graphql-api/interfaces"
	"github.com/Naheed-Rayhan/graphql-api/usecases"
)

func main() {
	// Initialize database
	config.InitDB()



	// Initialize repository
	courseRepo := database.NewCourseRepository(config.DB)
	
	// Initialize use case
	courseUseCase := usecases.NewCourseUseCase(courseRepo)

	// Initialize handler
	courseHandler := interfaces.NewCourseHandler(courseUseCase)

	// Setup router
	router := infrastructure.SetupRouter(courseHandler)

	// Start server
	router.Run(":8080")
}
