package Models

import "gorm.io/gorm"

type DroneLoad struct {
	gorm.Model
	DroneId             uint `json:"drone_id"`
	TotalLoadWeight     int
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	DroneLoadMedication []DroneLoadMedication
	DroneLoadLogs       []DroneLoadLogs
}
