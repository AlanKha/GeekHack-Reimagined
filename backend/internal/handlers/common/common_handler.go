package common

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
)

// Handler holds the datastore
type Handler struct {
	DB database.Datastore
}

// NewHandler creates a new handler
func NewHandler(db database.Datastore) *Handler {
	return &Handler{DB: db}
}
