package Controllers

import (
	"GoGin/Application"
	"GoGin/Models"
	Resources "GoGin/Resources/Medications"
	Validation "GoGin/Validations/Medications"

	"github.com/gin-gonic/gin"
	// To generate random file names
)

// Register new medication
func RegisterMedication(c *gin.Context) {
	r := Application.NewRequest(c)

	// Binding Request Body to drone Model
	var medication Models.Medication

	r.Context.ShouldBindJSON(&medication)

	// Validate Medication Model
	if r.ValidateRequest(Validation.InsertValidation(medication)).FailsValidation() {
		return
	}

	r.DB.Create(&medication)
	r.Created(Resources.MedicationResource(medication))
}
