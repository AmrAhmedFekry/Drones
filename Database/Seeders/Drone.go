package Seeders

import (
	"gorm.io/gorm"
)

// Use go faker to seed data
func seedDrones(DB *gorm.DB) {
	for i := 0; i < 5; i++ {
		// DB.Create(&Models.Drone{
		// 	SerialNumber: "123456789",
		// 	DroneModel
		// })
	}
}
