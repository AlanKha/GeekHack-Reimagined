package notification

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/common"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Handler holds the datastore
type Handler struct {
	*common.Handler
}

// NewHandler creates a new handler
func NewHandler(db database.Datastore) *Handler {
	return &Handler{common.NewHandler(db)}
}

// GetUserNotifications returns paginated notifications for the authenticated user
func (h *Handler) GetUserNotifications(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

	// For now, return empty array as placeholder
	notifications := []models.NotificationSummary{}
	total := int64(0)

	response := models.PaginatedResponse{
		Data:       notifications,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int((total + int64(perPage) - 1) / int64(perPage)),
		HasNext:    int64(page*perPage) < total,
		HasPrev:    page > 1,
	}

	c.JSON(http.StatusOK, response)
}

// MarkNotificationAsRead marks a specific notification as read
func (h *Handler) MarkNotificationAsRead(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid notification ID")
		return
	}

	_, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// This would need to be implemented in the datastore
	// err = h.DB.MarkNotificationAsRead(uint(notificationID), user.(*models.User).ID)

	// For now, just return success
	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

// MarkAllNotificationsAsRead marks all notifications as read for the authenticated user
func (h *Handler) MarkAllNotificationsAsRead(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// This would need to be implemented in the datastore
	c.JSON(http.StatusOK, gin.H{"message": "All notifications marked as read"})
}

// SubscribeToThread subscribes the user to thread notifications
func (h *Handler) SubscribeToThread(c *gin.Context) {
	threadID, err := strconv.ParseUint(c.Param("threadId"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid thread ID")
		return
	}

	user, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	_ = models.ThreadSubscription{
		UserID:   user.(*models.User).ID,
		ThreadID: uint(threadID),
		IsActive: true,
	}

	// This would need to be implemented in the datastore
	c.JSON(http.StatusOK, gin.H{"message": "Subscribed to thread notifications"})
}

// UnsubscribeFromThread unsubscribes the user from thread notifications
func (h *Handler) UnsubscribeFromThread(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("threadId"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid thread ID")
		return
	}

	_, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// This would need to be implemented in the datastore
	c.JSON(http.StatusOK, gin.H{"message": "Unsubscribed from thread notifications"})
}
