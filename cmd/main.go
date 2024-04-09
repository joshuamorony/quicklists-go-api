package main

import (
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// just store checklists in memory
// will be lost when server is restarted
var (
	checklists     = []Checklist{}
	checklistItems = []ChecklistItem{}
)

var sm *scs.SessionManager

func main() {
	sm = scs.New()
	sm.Cookie.Secure = true
	sm.Cookie.HttpOnly = true
	sm.Lifetime = 24 * time.Hour

	router := gin.Default()

	config := cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))
	router.Use(Adapt())

	router.POST("/login", login)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(Authorize())
	{
		protectedRoutes.POST("/logout", logout)

		protectedRoutes.GET("/checklists", getChecklists)
		protectedRoutes.POST("/checklists", addChecklist)

		protectedRoutes.GET("/checklists/:id", getChecklistByID)
		protectedRoutes.PATCH("/checklists/:id", editChecklistByID)
		protectedRoutes.DELETE("/checklists/:id", removeChecklistByID)
		protectedRoutes.POST("/checklists/:id/reset", resetChecklistByID)

		protectedRoutes.GET("/checklist-items", getChecklistItems)
		protectedRoutes.POST("/checklist-items/:id", addChecklistItem)
		protectedRoutes.PATCH("/checklist-items/:id", editChecklistItemByID)
		protectedRoutes.DELETE("/checklist-items/:id", removeChecklistItemByID)
	}

	router.Run("localhost:8080")
}
