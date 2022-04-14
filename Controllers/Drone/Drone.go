package Controllers

import (
	"GoGin/Application"
	"GoGin/Models"
	Resources "GoGin/Resources/Drones"

	"github.com/gin-gonic/gin"
)

// Register new drone
func RegisterDrone(c *gin.Context) {
	r := Application.NewRequest(c)

	// Binding Request Body to drone Model
	var drone Models.Drone
	r.Context.ShouldBindJSON(&drone)
	// Validate user Model
	// if r.ValidateRequest(Sellers.InsertValidation(product)).FailsValidation() {
	// 	return
	// }
	r.DB.Create(&drone)
	r.Created(Resources.DroneResource(drone))
}
