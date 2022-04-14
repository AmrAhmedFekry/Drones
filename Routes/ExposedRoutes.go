package Routes

import Controllers "GoGin/Controllers/Drone"

func (app RouterApp) DroneRoutes() {
	app.Gin.POST("/api/drone/register", Controllers.RegisterDrone)
}
