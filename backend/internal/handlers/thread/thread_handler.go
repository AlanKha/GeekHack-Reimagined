package thread

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/common"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Handler holds the datastore
type Handler struct {
	*common.Handler
}

// NewHandler creates a new handler
func NewHandler(db database.Datastore) *Handler {
	return &Handler{common.NewHandler(db)}
}

func (h *Handler) CreateThread(c *gin.Context) {
	var body struct {
		Title      string
		Content    string
		CategoryID uint
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	user, _ := c.Get("user")
	now := time.Now()

	thread := models.Thread{
		Title:        body.Title,
		Content:      body.Content,
		UserID:       user.(*models.User).ID,
		CategoryID:   body.CategoryID,
		PostCount:    1, // Original post counts as 1
		ReplyCount:   0,
		LastActivity: now,
		IsActive:     true,
	}

	err := h.DB.CreateThread(&thread)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create thread")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Thread created", "thread": thread.ToSummary()})
}

func (h *Handler) GetThreads(c *gin.Context) {
	threads, err := h.DB.GetThreads()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get threads")
		return
	}

	// Convert to DTOs for optimized response
	summaries := make([]models.ThreadSummary, len(threads))
	for i, thread := range threads {
		summaries[i] = thread.ToSummary()
	}

	c.JSON(http.StatusOK, gin.H{"threads": summaries})
}

func (h *Handler) GetThread(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid thread ID")
		return
	}

	thread, err := h.DB.GetThreadByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Thread not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"thread": thread})
}

func (h *Handler) UpdateThread(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid thread ID")
		return
	}

	var body struct {
		Title   string
		Content string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	thread, err := h.DB.GetThreadByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Thread not found")
		return
	}

	user, _ := c.Get("user")
	if thread.UserID != user.(*models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	thread.Title = body.Title
	thread.Content = body.Content

	err = h.DB.UpdateThread(thread)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update thread")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thread updated", "thread": thread})
}

func (h *Handler) DeleteThread(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid thread ID")
		return
	}

	thread, err := h.DB.GetThreadByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Thread not found")
		return
	}

	user, _ := c.Get("user")
	if thread.UserID != user.(*models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err = h.DB.DeleteThread(thread)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete thread")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thread deleted"})
}
