package main

import (
	"GoGin/Application"
	"GoGin/Routes"
	"GoGin/Utils"
)

func main() {
	// Initialize the application with all required dependencies and configurations and run the application
	app := Application.NewApp()

	// migrate models
	app.Migrate()
	// Seed data
	app.Seed()

	// Close Application connection
	Application.CloseConnection(app)
	// Get application routes
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	// Run your own schedule Functions
	Utils.RunScheduleTasks()
	// Run application
	app.Gin.Run(":8080")
}
