package auth

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/common"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
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

func (h *Handler) Register(c *gin.Context) {
	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
		JoinedAt: time.Now(),
		LastSeen: time.Now(),
	}
	err = h.DB.CreateUser(&user)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Username already exists")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (h *Handler) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	user, err := h.DB.GetUserByEmail(body.Email)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid username or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid username or password")
		return
	}

	// Update last seen time
	user.LastSeen = time.Now()
	h.DB.UpdateUser(user) // You'll need to add this method to datastore

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create token")
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
