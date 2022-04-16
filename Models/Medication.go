package Models

import "gorm.io/gorm"

type Medication struct {
	gorm.Model
	Code                string `json:"code" gorm:"type:varchar(100) not null"`
	Name                string `json:"name" gorm:"type:varchar(100) not null"`
	Weight              int    `json:"weight"`
	Image               string `json:"image"`
	DroneLoadMedication []DroneLoadMedication
}
