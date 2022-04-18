package Models

import "gorm.io/gorm"

type DroneBatteryLogs struct {
	gorm.Model
	DroneId              uint `json:"drone_id"`
	DroneBatteryCapacity int  `json:"drone_battery"`
}
