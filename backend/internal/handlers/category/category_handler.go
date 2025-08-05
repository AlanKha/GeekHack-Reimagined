package category

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

func (h *Handler) CreateCategory(c *gin.Context) {
	var body struct {
		Name         string
		Description  string
		Slug         string
		DisplayOrder int
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	category := models.Category{
		Name:         body.Name,
		Description:  body.Description,
		Slug:         body.Slug,
		DisplayOrder: body.DisplayOrder,
		IsActive:     true, // Default to active
	}

	err := h.DB.CreateCategory(&category)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create category")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "category": category.ToSummary()})
}

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.DB.GetCategories()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get categories")
		return
	}

	// Convert to DTOs for optimized response
	summaries := make([]models.CategorySummary, len(categories))
	for i, cat := range categories {
		summaries[i] = cat.ToSummary()
	}

	c.JSON(http.StatusOK, gin.H{"categories": summaries})
}

func (h *Handler) GetCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.DB.GetCategoryByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Category not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category.ToSummary()})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	var body struct {
		Name        string
		Description string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	category, err := h.DB.GetCategoryByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Category not found")
		return
	}

	category.Name = body.Name
	category.Description = body.Description

	err = h.DB.UpdateCategory(category)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update category")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated", "category": category})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.DB.GetCategoryByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Category not found")
		return
	}

	err = h.DB.DeleteCategory(category)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
