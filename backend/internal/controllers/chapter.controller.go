package controllers

import (
	"backend/db"
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CreateChapter(c *gin.Context) {
	var chapter models.Chapter

	if err := c.ShouldBindJSON(&chapter); err != nil {
		log.Print("Missing data to create a chapter: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&chapter).Error; err != nil {
		log.Print("Error creating chapter in db: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, chapter)
}

func GetChapterById(c *gin.Context) {
	id := c.Param("id")
	var chapter models.Chapter

	if err := db.DB.First(&chapter, id).Error; err != nil {
		log.Print("Chapter not found with id: ", id, " Error: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}

	c.JSON(http.StatusOK, chapter)
}

func DeleteChapter(c *gin.Context) {
	id := c.Param("id")
	var chapter models.Chapter

	// First check if the chapter exists
	if err := db.DB.First(&chapter, id).Error; err != nil {
		log.Print("Chapter not found with id: ", id, " Error: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}

	// Delete the chapter
	if err := db.DB.Delete(&chapter).Error; err != nil {
		log.Print("Error deleting chapter with id: ", id, " Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chapter deleted successfully"})
}

func UpdateChapter(c *gin.Context) {
	id := c.Param("id")
	var chapter models.Chapter

	// First check if the chapter exists
	if err := db.DB.First(&chapter, id).Error; err != nil {
		log.Print("Chapter not found with id: ", id, " Error: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}

	// Bind the update data from request body
	var updateData models.Chapter
	if err := c.ShouldBindJSON(&updateData); err != nil {
		log.Print("Invalid update data for chapter: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the chapter
	if err := db.DB.Model(&chapter).Updates(updateData).Error; err != nil {
		log.Print("Error updating chapter with id: ", id, " Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chapter)
}
