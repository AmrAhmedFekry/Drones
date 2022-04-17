package Controllers

import (
	"GoGin/Application"
	"GoGin/Models"

	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
)

// Load medication request data
type LoadMedicationRequest struct {
	DroneId       int     `json:"drone_id"`
	MedicationIds []int   `json:"medication_items"`
	Latitude      float64 `json:"Latitude"`
	Longitude     float64 `json:"Longitude"`
}

// Load medication request data
type ChangeDroneLoadStateRequest struct {
	DroneLoadId    int               `json:"drone_load_id"`
	DroneLoadState Models.DroneState `json:"drone_load_state"`
	DroneBattery   int               `json:"drone_battery"`
}

// Change drone load state
func ChangeDroneLoadState(c *gin.Context) {
	r := Application.NewRequest(c)
	var requestData ChangeDroneLoadStateRequest
	r.Context.ShouldBindJSON(&requestData)

	// Get drone load by id
	var droneLoad Models.DroneLoad
	if err := r.DB.Where("id = ?", requestData.DroneLoadId).First(&droneLoad).Error; err != nil {
		r.ResourceNotFound("drone_load")
		return
	}
	// Update drone load state
	droneLoad.DroneLoadState = requestData.DroneLoadState
	r.DB.Save(&droneLoad)

	// Update drone state
	var drone Models.Drone
	if err := r.DB.Where("id = ?", droneLoad.DroneId).First(&drone).Error; err != nil {
		r.ResourceNotFound("drone")
		return
	}
	drone.DroneState = requestData.DroneLoadState
	drone.BatteryCapacity = requestData.DroneBattery
	r.DB.Save(&drone)

	// Create drone load logs
	var droneLoadLogs Models.DroneLoadLogs
	droneLoadLogs.DroneLoadId = droneLoad.ID
	droneLoadLogs.DroneState = requestData.DroneLoadState
	droneLoadLogs.DroneBatteryCapacity = drone.BatteryCapacity
	r.DB.Create(&droneLoadLogs)

	r.Success("Drone Load State Changed")
}

// Load medication request data
func LoadMedications(c *gin.Context) {
	r := Application.NewRequest(c)
	var requestData LoadMedicationRequest
	r.Context.ShouldBindJSON(&requestData)

	// Get drone by id
	var drone Models.Drone
	if err := r.DB.Where("id = ?", requestData.DroneId).First(&drone).Error; err != nil {
		r.ResourceNotFound("drone")
		return
	}

	// Basic validation about drone state and battery capacity
	if drone.DroneState != "IDLE" {
		r.BadRequest(gotrans.T("drone_is_not_idle"))
		return
	}
	if drone.BatteryCapacity < 25 {
		r.BadRequest(gotrans.T("drone_battery_is_low"))
		return
	}

	// Get medications by ids
	var medications []Models.Medication
	r.DB.Where("id IN (?)", requestData.MedicationIds).Find(&medications)

	// Check if medications wight is less than drone wight limit
	totalWeight := 0
	for _, medication := range medications {
		totalWeight += medication.Weight
	}
	if totalWeight > drone.WeightLimit {
		r.BadRequest(gotrans.T("medication_weight_is_over_drone_weight_limit"))
		return
	}

	// Update drone state and battery capacity
	drone.DroneState = "LOADING"
	r.DB.Save(&drone)

	// Create new drone load
	var droneLoad Models.DroneLoad
	droneLoad.DroneId = drone.ID
	droneLoad.TotalLoadWeight = totalWeight
	droneLoad.DroneLoadState = "LOADING"
	droneLoad.Latitude = requestData.Latitude
	droneLoad.Longitude = requestData.Longitude
	r.DB.Create(&droneLoad)

	// Create drone load medications
	var droneLoadMedications []Models.DroneLoadMedication
	for _, medicationId := range requestData.MedicationIds {
		droneLoadMedications = append(droneLoadMedications, Models.DroneLoadMedication{
			DroneLoadId:  droneLoad.ID,
			MedicationId: uint(medicationId),
		})
	}
	r.DB.Create(&droneLoadMedications)

	// Create drone load logs
	var droneLoadLogs Models.DroneLoadLogs
	droneLoadLogs.DroneLoadId = droneLoad.ID
	droneLoadLogs.DroneState = "LOADING"
	droneLoadLogs.DroneBatteryCapacity = drone.BatteryCapacity
	r.DB.Create(&droneLoadLogs)
	r.Success(medications)
}
