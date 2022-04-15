package Medications

import (
	"GoGin/Models"
	"os"
)

/*
	The Idea behind this snippet is to use this function to return only required fields
	of every model in case of single medication, list of medications and depend on the API requested Actor
*/

// Map of  medication data in case of single medication and if there are common fields
// Between singel and list of medications then use this function to set them
func MedicationResource(medication Models.Medication) map[string]interface{} {
	medicationResource := make(map[string]interface{})
	medicationResource["id"] = medication.ID
	medicationResource["Name"] = medication.Name
	medicationResource["code"] = medication.Code
	medicationResource["Weight"] = medication.Weight
	medicationResource["Image"] = os.Getenv("APP_URL") + "/" + medication.Image

	return medicationResource
}

// Map of medications data
func medicationsResource(medications []Models.Medication) []map[string]interface{} {
	mappedMedications := make([]map[string]interface{}, 0)
	for _, medication := range medications {
		mappedMedications = append(mappedMedications, MedicationResource(medication))
	}
	return mappedMedications
}
