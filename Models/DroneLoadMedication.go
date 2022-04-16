package Models

import "gorm.io/gorm"

type DroneLoadMedication struct {
	gorm.Model
	DroneLoadId  uint `json:"drone_load_id"`
	MedicationId uint `json:"medication_id"`
}
