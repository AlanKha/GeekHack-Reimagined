package reaction

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

func (h *Handler) CreateReaction(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var body struct {
		ReactionType string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	user, _ := c.Get("user")

	reaction := models.Reaction{
		PostID:       uint(postID),
		UserID:       user.(*models.User).ID,
		ReactionType: models.ReactionType(body.ReactionType),
	}

	err = h.DB.CreateReaction(&reaction)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create reaction")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reaction created", "reaction": reaction})
}

func (h *Handler) GetReactions(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	reactions, err := h.DB.GetReactionsByPostID(uint(postID))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get reactions")
		return
	}

	c.JSON(http.StatusOK, gin.H{"reactions": reactions})
}
