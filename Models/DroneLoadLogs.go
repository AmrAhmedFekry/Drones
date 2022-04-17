package Models

import "gorm.io/gorm"

type DroneLoadLogs struct {
	gorm.Model
	DroneLoadId          uint       `json:"drone_load_id"`
	DroneState           DroneState `json:"drone_state"`
	DroneBatteryCapacity int        `json:"drone_battery"`
}
