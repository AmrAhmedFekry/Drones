package Routes

import (
	DroneController "GoGin/Controllers/Drone"
	FileController "GoGin/Controllers/File"
	MedicationController "GoGin/Controllers/Medication"
)

func (app RouterApp) DroneRoutes() {
	app.Gin.POST("/api/drone/register", DroneController.RegisterDrone)
	app.Gin.GET("/api/available/loading/drones", DroneController.ListOfAvailableDronesForLoading)
}

func (app RouterApp) MedicationRoutes() {
	app.Gin.POST("/api/medication", MedicationController.RegisterMedication)
}

func (app RouterApp) MediaRoutes() {
	app.Gin.POST("/api/upload", FileController.FileUpload)
	app.Gin.GET(":file", FileController.ShowFile)
}
