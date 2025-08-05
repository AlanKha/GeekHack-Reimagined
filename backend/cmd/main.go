package main

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/auth"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/category"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/moderation_log"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/post"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/reaction"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/thread"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/user"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := database.NewDBClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	authHandler := auth.NewHandler(db)
	m := middleware.NewMiddleware(db)

	r := gin.Default()

	r.POST("/api/register", authHandler.Register)
	r.POST("/api/login", authHandler.Login)

	r.GET("/api/validate", m.RequireAuth, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "I'm logged in"})
	})

	// Thread Routes
	threadHandler := thread.NewHandler(db)
	r.POST("/api/threads", m.RequireAuth, threadHandler.CreateThread)
	r.GET("/api/threads", threadHandler.GetThreads)
	r.GET("/api/threads/:id", threadHandler.GetThread)
	r.PUT("/api/threads/:id", m.RequireAuth, threadHandler.UpdateThread)
	r.DELETE("/api/threads/:id", m.RequireAuth, threadHandler.DeleteThread)

	// Post Routes
	postHandler := post.NewHandler(db)
	r.POST("/api/threads/:id/posts", m.RequireAuth, postHandler.CreatePost)
	r.GET("api/posts/:id", postHandler.GetPost)
	r.PUT("/api/posts/:id", m.RequireAuth, postHandler.UpdatePost)
	r.DELETE("/api/posts/:id", m.RequireAuth, postHandler.DeletePost)

	// User Routes
	userHandler := user.NewHandler(db)
	r.GET("/api/users/:id", userHandler.GetUser)
	r.PUT("/api/users/:id", m.RequireAuth, userHandler.UpdateUser)

	// Category Routes
	categoryHandler := category.NewHandler(db)
	r.POST("/api/categories", m.RequireAuth, categoryHandler.CreateCategory)
	r.GET("/api/categories", categoryHandler.GetCategories)
	r.GET("/api/categories/:id", categoryHandler.GetCategory)
	r.PUT("/api/categories/:id", m.RequireAuth, categoryHandler.UpdateCategory)
	r.DELETE("/api/categories/:id", m.RequireAuth, categoryHandler.DeleteCategory)

	// Reaction Routes
	reactionHandler := reaction.NewHandler(db)
	r.POST("/api/posts/:id/reactions", m.RequireAuth, reactionHandler.CreateReaction)
	r.GET("/api/posts/:id/reactions", reactionHandler.GetReactions)

	// Moderation Log Routes
	moderationLogHandler := moderation_log.NewHandler(db)
	r.POST("/api/moderation-logs", m.RequireAuth, moderationLogHandler.CreateModerationLog)
	r.GET("/api/moderation-logs", m.RequireAuth, moderationLogHandler.GetModerationLogs)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
