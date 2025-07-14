
package handlers

import (
	"fmt"
	"net/http"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context, db *gorm.DB) {
	threadID := c.Param("id")

	var body struct {
		Content string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	user, _ := c.Get("user")

	post := models.Post{
		Content:  body.Content,
		UserID:   user.(models.User).ID,
		ThreadID: parseUint(threadID),
	}

	result := db.Create(&post)

	if result.Error != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create post")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": post})
}

func GetPost(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var post models.Post
	result := db.Preload("User").First(&post, id)

	if result.Error != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func UpdatePost(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var body struct {
		Content string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	var post models.Post
	db.First(&post, id)

	if post.ID == 0 {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	user, _ := c.Get("user")
	if post.UserID != user.(models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	db.Model(&post).Updates(models.Post{Content: body.Content})

	c.JSON(http.StatusOK, gin.H{"message": "Post updated", "post": post})
}

func DeletePost(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var post models.Post
	db.First(&post, id)

	if post.ID == 0 {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	user, _ := c.Get("user")
	if post.UserID != user.(models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	db.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func parseUint(s string) uint {
	var i uint64
	fmt.Sscanf(s, "%d", &i)
	return uint(i)
}
