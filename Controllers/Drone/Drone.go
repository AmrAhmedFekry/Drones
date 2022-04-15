package Controllers

import (
	"GoGin/Application"
	"GoGin/Models"
	Resources "GoGin/Resources/Drones"
	Validation "GoGin/Validations/Drones"

	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
)

// Register new drone
func RegisterDrone(c *gin.Context) {
	r := Application.NewRequest(c)

	// Binding Request Body to drone Model
	var drone Models.Drone
	// Set Initial value.
	drone.BatteryCapacity = 100
	drone.DroneState = "IDLE"

	r.Context.ShouldBindJSON(&drone)

	// Validate if serial number is unique
	if r.DB.Where("serial_number = ?", drone.SerialNumber).Find(&Models.Drone{}).RowsAffected > 0 {
		r.BadRequest(gotrans.T("serial_number_already_exist"))
		return
	}
	// Validate Drone Model
	if r.ValidateRequest(Validation.InsertValidation(drone)).FailsValidation() {
		return
	}

	r.DB.Create(&drone)
	r.Created(Resources.DroneResource(drone))
}

// List with all available drones for loading
func ListOfAvailableDronesForLoading(c *gin.Context) {

}
