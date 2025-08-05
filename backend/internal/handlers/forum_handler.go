package handlers

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ForumHandler provides aggregate forum data
type ForumHandler struct {
	DB database.Datastore
}

// NewForumHandler creates a new forum handler
func NewForumHandler(db database.Datastore) *ForumHandler {
	return &ForumHandler{DB: db}
}

// GetForumStats provides forum-wide statistics
func (h *ForumHandler) GetForumStats(c *gin.Context) {
	// This would need to be implemented in the datastore
	// For now, we'll return a placeholder
	stats := models.ForumStats{
		TotalUsers:    1000,
		TotalThreads:  150,
		TotalPosts:    2500,
		OnlineUsers:   25,
		NewestUser:    "john_doe",
		PopularThread: "Welcome to GeekHack!",
	}

	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

// GetCategoriesWithStats returns categories with their thread/post counts
func (h *ForumHandler) GetCategoriesWithStats(c *gin.Context) {
	categories, err := h.DB.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get categories"})
		return
	}

	// Convert to summaries which include the denormalized counts
	summaries := make([]models.CategorySummary, len(categories))
	for i, cat := range categories {
		summaries[i] = cat.ToSummary()
	}

	c.JSON(http.StatusOK, gin.H{"categories": summaries})
}

// GetThreadsByCategory returns paginated threads for a category
func (h *ForumHandler) GetThreadsByCategory(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("categoryId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

	// This would need pagination logic in the datastore
	// For now, return all threads (this should be implemented in datastore)
	threads, err := h.DB.GetThreads() // This needs to be updated to filter by category
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get threads"})
		return
	}

	// Convert to summaries
	summaries := make([]models.ThreadSummary, len(threads))
	for i, thread := range threads {
		summaries[i] = thread.ToSummary()
	}

	// Create paginated response
	response := models.PaginatedResponse{
		Data:       summaries,
		Page:       page,
		PerPage:    perPage,
		Total:      int64(len(summaries)),
		TotalPages: (len(summaries) + perPage - 1) / perPage,
		HasNext:    page*perPage < len(summaries),
		HasPrev:    page > 1,
	}

	c.JSON(http.StatusOK, response)
}
