package main

import (
	"log"
	"net/http"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDBClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	h := handlers.NewHandler(db)
	m := middleware.NewMiddleware(db)

	r := gin.Default()

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	r.GET("/validate", m.RequireAuth, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "I'm logged in"})
	})

	// Thread Routes
	r.POST("/api/threads", m.RequireAuth, h.CreateThread)
	r.GET("/api/threads", h.GetThreads)
	r.GET("/api/threads/:id", h.GetThread)
	r.PUT("/api/threads/:id", m.RequireAuth, h.UpdateThread)
	r.DELETE("/api/threads/:id", m.RequireAuth, h.DeleteThread)

	// Post Routes
	r.POST("/api/threads/:id/posts", m.RequireAuth, h.CreatePost)
	r.GET("api/posts/:id", h.GetPost)
	r.PUT("/api/posts/:id", m.RequireAuth, h.UpdatePost)
	r.DELETE("/api/posts/:id", m.RequireAuth, h.DeletePost)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
