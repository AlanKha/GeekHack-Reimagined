
package handlers

import (
	"net/http"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateThread(c *gin.Context, db *gorm.DB) {
	var body struct {
		Title   string
		Content string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	user, _ := c.Get("user")

	thread := models.Thread{
		Title:   body.Title,
		Content: body.Content,
		UserID:  user.(models.User).ID,
	}

	result := db.Create(&thread)

	if result.Error != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create thread")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Thread created", "thread": thread})
}

func GetThreads(c *gin.Context, db *gorm.DB) {
	var threads []models.Thread
	db.Preload("User").Find(&threads)

	c.JSON(http.StatusOK, gin.H{"threads": threads})
}

func GetThread(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var thread models.Thread
	result := db.Preload("User").Preload("Posts.User").First(&thread, id)

	if result.Error != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Thread not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"thread": thread})
}

func UpdateThread(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var body struct {
		Title   string
		Content string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	var thread models.Thread
	db.First(&thread, id)

	if thread.ID == 0 {
		utils.RespondWithError(c, http.StatusNotFound, "Thread not found")
		return
	}

	user, _ := c.Get("user")
	if thread.UserID != user.(models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	db.Model(&thread).Updates(models.Thread{Title: body.Title, Content: body.Content})

	c.JSON(http.StatusOK, gin.H{"message": "Thread updated", "thread": thread})
}

func DeleteThread(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var thread models.Thread
	db.First(&thread, id)

	if thread.ID == 0 {
		utils.RespondWithError(c, http.StatusNotFound, "Thread not found")
		return
	}

	user, _ := c.Get("user")
	if thread.UserID != user.(models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	db.Delete(&thread)

	c.JSON(http.StatusOK, gin.H{"message": "Thread deleted"})
}
