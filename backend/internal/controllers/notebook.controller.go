package controllers

import (
	"backend/db"
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetNotebookById(c *gin.Context) {
	var notebook models.Notebook
	id := c.Param("id")

	if err := db.DB.First(&notebook, id).Error; err != nil {
		log.Print("Notebook not found with id: ", id, " Error: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Notebook not found"})
		return
	}

	c.JSON(http.StatusOK, notebook)
}

func CreateNotebook(c *gin.Context) {
	var notebook models.Notebook

	if err := c.ShouldBindJSON(&notebook); err != nil {
		log.Print("Missing data to create a notebook", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&notebook).Error; err != nil {
		log.Print("Error creating Notebook in db: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, notebook)
}

func DeleteNotebook(c *gin.Context) {
	var notebook models.Notebook
	id := c.Param("id")

	if err := db.DB.First(&notebook, id).Error; err != nil {
		log.Print("Notebook not found with id: ", id, " Error: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Notebook not found"})
		return
	}

	if err := db.DB.Delete(&notebook, id).Error; err != nil {
		log.Print("Error deleting Notebook with id: ", id, "Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notebook deleted successfully"})
}

func UpdateNotebook(c *gin.Context) {
	id := c.Param("id")
	var notebook models.Notebook

	// First check if the notebook exists
	if err := db.DB.First(&notebook, id).Error; err != nil {
		log.Print("Notebook not found with id: ", id, " Error: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Notebook not found"})
		return
	}

	// Bind the update data from request body
	var updateData models.Notebook
	if err := c.ShouldBindJSON(&updateData); err != nil {
		log.Print("Invalid update data for notebook: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the notebook
	if err := db.DB.Model(&notebook).Updates(updateData).Error; err != nil {
		log.Print("Error updating Notebook with id: ", id, " Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notebook)
}
