package Drones

import (
	"GoGin/Models"
)

/*
	The Idea behind this snippet is to use this function to return only required fields
	of every model in case of single drone, list of drones and depend on the API requested Actor
*/

// Map of  drone data in case of single drone and if there are common fields
// Between singel and list of drones then use this function to set them
func DroneResource(drone Models.Drone) map[string]interface{} {
	DroneResource := make(map[string]interface{})
	DroneResource["id"] = drone.ID
	DroneResource["SerialNumber"] = drone.SerialNumber
	DroneResource["DroneModel"] = drone.DroneModel
	DroneResource["DroneState"] = drone.DroneState
	DroneResource["WeightLimit"] = drone.WeightLimit
	DroneResource["BatteryCapacity"] = drone.BatteryCapacity

	return DroneResource
}

// Map of drones data
func DronesResource(drones []Models.Drone) []map[string]interface{} {
	mappedDrones := make([]map[string]interface{}, 0)
	for _, drone := range drones {
		mappedDrones = append(mappedDrones, DroneResource(drone))
	}
	return mappedDrones
}
