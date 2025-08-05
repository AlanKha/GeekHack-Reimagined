package moderation_log

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/common"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler holds the datastore
type Handler struct {
	*common.Handler
}

// NewHandler creates a new handler
func NewHandler(db database.Datastore) *Handler {
	return &Handler{common.NewHandler(db)}
}

func (h *Handler) CreateModerationLog(c *gin.Context) {
	var body struct {
		UserID   uint
		Action   string
		Reason   string
		ThreadID uint
		PostID   uint
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	var threadID *uint
	if body.ThreadID != 0 {
		threadID = &body.ThreadID
	}

	var postID *uint
	if body.PostID != 0 {
		postID = &body.PostID
	}

	moderationLog := models.ModerationLog{
		ModeratorID: body.UserID, // Assuming this is the moderator
		Action:      models.ModerationAction(body.Action),
		Reason:      body.Reason,
		ThreadID:    threadID,
		PostID:      postID,
	}

	err := h.DB.CreateModerationLog(&moderationLog)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create moderation log")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Moderation log created", "moderation_log": moderationLog})
}

func (h *Handler) GetModerationLogs(c *gin.Context) {
	moderationLogs, err := h.DB.GetModerationLogs()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get moderation logs")
		return
	}

	c.JSON(http.StatusOK, gin.H{"moderation_logs": moderationLogs})
}
