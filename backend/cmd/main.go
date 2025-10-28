package main

import (
	"backend/db"
	"backend/internal/auth"
	"backend/internal/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db.InitDB()

	// Initialize authentication
	auth.InitAuth()

	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Cookie"},
		ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	}))

	// Auth routes
	r.GET("/auth/:provider", auth.BeginAuth)
	r.GET("/auth/:provider/callback", auth.AuthCallback)
	r.GET("/auth/user", auth.GetCurrentUser)
	r.GET("/logout/:provider", auth.Logout)

	// Notebook routes
	r.POST("/notebook", controllers.CreateNotebook)
	r.GET("/notebook/:id", controllers.GetNotebookById)
	r.PUT("/notebook/:id", controllers.UpdateNotebook)
	r.DELETE("/notebook/:id", controllers.DeleteNotebook)

	// Chapter routes
	r.POST("/chapter", controllers.CreateChapter)
	r.GET("/chapter/:id", controllers.GetChapterById)
	r.PUT("/chapter/:id", controllers.UpdateChapter)
	r.DELETE("/chapter/:id", controllers.DeleteChapter)

	r.Run(":8080")
}
