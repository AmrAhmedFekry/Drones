package Routes

import "GoGin/Application"

type RouterApp struct {
	*Application.Bootstrap
}

// Set your routes here
func (app RouterApp) Routing() {
	app.DroneRoutes()
	app.MedicationRoutes()
	app.MediaRoutes()
	app.LoadingRoutes()
}
