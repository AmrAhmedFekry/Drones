package Drones

import (
	"GoGin/Models"
	"GoGin/Validations"

	validation "github.com/go-ozzo/ozzo-validation"
)

func InsertValidation(drone Models.Drone) validation.Errors {
	return validation.Errors{
		"serial_number": validation.Validate(drone.SerialNumber, Validations.RequiredRule(), Validations.MaxCharacter(100)),
		"model":         validation.Validate(drone.DroneModel, Validations.RequiredRule(), Validations.EnumValues("model", "Lightweight", "Middleweight", "Cruiserweight", "Heavyweight")),
		"weight_limit":  validation.Validate(drone.WeightLimit, Validations.RequiredRule(), Validations.MaxValue(500)),
	}
}
