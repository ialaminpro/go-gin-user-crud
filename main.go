package main

import (
	"myginapp/controllers"
	"myginapp/models"
	"myginapp/repositories"
	"myginapp/routes"
	"myginapp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	router := gin.Default()

	// Serve the static files in the "static" directory
	router.Static("/static", "./static")

	// Set up the routes
	routes.UserRoutes(router, userController)

	router.Run(":8080")
}
