package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	threadID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid thread ID")
		return
	}

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
		UserID:   user.(*models.User).ID,
		ThreadID: uint(threadID),
	}

	err = h.DB.CreatePost(&post)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create post")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": post})
}

func (h *Handler) GetPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := h.DB.GetPostByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var body struct {
		Content string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	post, err := h.DB.GetPostByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	user, _ := c.Get("user")
	if post.UserID != user.(*models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	post.Content = body.Content
	err = h.DB.UpdatePost(post)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update post")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated", "post": post})
}

func (h *Handler) DeletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := h.DB.GetPostByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	user, _ := c.Get("user")
	if post.UserID != user.(*models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err = h.DB.DeletePost(post)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete post")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func parseUint(s string) uint {
	var i uint64
	fmt.Sscanf(s, "%d", &i)
	return uint(i)
}
