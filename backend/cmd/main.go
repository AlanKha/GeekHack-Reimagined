
package main

import (
	"net/http"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/register", func(c *gin.Context) { handlers.Register(c, database.DB) })
	r.POST("/login", func(c *gin.Context) { handlers.Login(c, database.DB) })

	r.GET("/validate", middleware.RequireAuth, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "I'm logged in"})
	})

	// Thread Routes
	r.POST("/api/threads", middleware.RequireAuth, func(c *gin.Context) { handlers.CreateThread(c, database.DB) })
	r.GET("/api/threads", func(c *gin.Context) { handlers.GetThreads(c, database.DB) })
	r.GET("/api/threads/:id", func(c *gin.Context) { handlers.GetThread(c, database.DB) })
	r.PUT("/api/threads/:id", middleware.RequireAuth, func(c *gin.Context) { handlers.UpdateThread(c, database.DB) })
	r.DELETE("/api/threads/:id", middleware.RequireAuth, func(c *gin.Context) { handlers.DeleteThread(c, database.DB) })

	// Post Routes
	r.POST("/api/threads/:id/posts", middleware.RequireAuth, func(c *gin.Context) { handlers.CreatePost(c, database.DB) })
	r.GET("/api/posts/:id", func(c *gin.Context) { handlers.GetPost(c, database.DB) })
	r.PUT("/api/posts/:id", middleware.RequireAuth, func(c *gin.Context) { handlers.UpdatePost(c, database.DB) })
	r.DELETE("/api/posts/:id", middleware.RequireAuth, func(c *gin.Context) { handlers.DeletePost(c, database.DB) })

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
