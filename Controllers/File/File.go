package Controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Upload medication image
func FileUpload(c *gin.Context) {
	image, err := c.FormFile("image")

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(image.Filename)

	// Generate random image name for the new uploaded file so it doesn't override the old file with same name
	newImageName := uuid.New().String() + extension

	// Directory Path to save the image
	dirPath := "./Storage/" + newImageName
	// The file is received, so let's save it
	if err := c.SaveUploadedFile(image, dirPath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message":  "Your file has been successfully uploaded.",
		"fileName": newImageName,
		"filePath": c.Request.Host + "/" + newImageName,
	})
}

// Show uploaded file
func ShowFile(c *gin.Context) {
	fileName := c.Param("file")
	filePath := "./Storage/" + fileName

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "File not found",
		})
		return
	}

	// File exists, so let's send it
	c.File(filePath)
}
