package Models

import "gorm.io/gorm"

type DroneLoadLogs struct {
	gorm.Model
	DroneLoadId uint   `json:"drone_load_id"`
	DroneState  string `json:"drone_state"`
}
