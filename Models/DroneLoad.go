package Models

import "gorm.io/gorm"

type DroneLoad struct {
	gorm.Model
	DroneId             uint `json:"drone_id"`
	TotalLoadWeight     int
	DroneLoadState      DroneState `json:"drone_load_state"`
	Latitude            float64    `json:"latitude"`
	Longitude           float64    `json:"longitude"`
	DroneLoadMedication []DroneLoadMedication
	DroneLoadLogs       []DroneLoadLogs
}
