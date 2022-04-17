package Routes

import (
	DroneController "GoGin/Controllers/Drone"
	FileController "GoGin/Controllers/File"
	LoadingMedicationController "GoGin/Controllers/Loading"
	MedicationController "GoGin/Controllers/Medication"
)

// Available Routes for drones
func (app RouterApp) DroneRoutes() {
	app.Gin.POST("/api/drone/register", DroneController.RegisterDrone)
	app.Gin.GET("/api/available/loading/drones", DroneController.ListOfAvailableDronesForLoading)
	app.Gin.GET("/api/drone/battery_level/:droneId", DroneController.ShoWDroneBatteryCapacity)
}

func (app RouterApp) MedicationRoutes() {
	app.Gin.POST("/api/medication", MedicationController.RegisterMedication)
}

func (app RouterApp) MediaRoutes() {
	app.Gin.POST("/api/upload", FileController.FileUpload)
	app.Gin.GET(":file", FileController.ShowFile)
}

func (app RouterApp) LoadingRoutes() {
	app.Gin.POST("/api/load", LoadingMedicationController.LoadMedications)
	app.Gin.PUT("/api/change_drone_load_state", LoadingMedicationController.ChangeDroneLoadState)
}
