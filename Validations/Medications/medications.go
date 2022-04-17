package Medications

import (
	"GoGin/Models"
	"GoGin/Validations"

	validation "github.com/go-ozzo/ozzo-validation"
)

func InsertValidation(medication Models.Medication) validation.Errors {
	return validation.Errors{
		"name":   validation.Validate(medication.Name, Validations.RequiredRule(), Validations.MatchRegex("^[A-Za-z0-9_-]*$", "medication_name")),
		"weight": validation.Validate(medication.Weight, Validations.RequiredRule()),
		"code":   validation.Validate(medication.Code, Validations.RequiredRule(), Validations.MatchRegex("^[A-Z0-9_-]*$", "medication_code")),
	}
}
