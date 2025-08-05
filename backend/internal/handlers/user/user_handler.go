package user

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

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.DB.GetUserByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user.ToProfile()})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var body struct {
		AvatarURL string
		Signature string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	user, err := h.DB.GetUserByID(uint(id))

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found")
		return
	}

	currentUser, _ := c.Get("user")
	if user.ID != currentUser.(*models.User).ID {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user.AvatarURL = body.AvatarURL
	user.Signature = body.Signature

	err = h.DB.UpdateUser(user)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated", "user": user})
}
