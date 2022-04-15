package Medications

import (
	"GoGin/Models"
	"GoGin/Validations"

	validation "github.com/go-ozzo/ozzo-validation"
)

func InsertValidation(medication Models.Medication) validation.Errors {
	return validation.Errors{
		"name":   validation.Validate(medication.Name, Validations.RequiredRule()),
		"weight": validation.Validate(medication.Weight, Validations.RequiredRule()),
		"code":   validation.Validate(medication.Code, Validations.RequiredRule()),
	}
}
