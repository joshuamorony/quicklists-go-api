package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// just store checklists in memory
// will be lost when server is restarted
var (
	checklists     = []Checklist{}
	checklistItems = []ChecklistItem{}
)

func main() {
	router := gin.Default()

	config := cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))
	router.GET("/checklists", getChecklists)
	router.POST("/checklists", addChecklist)

	router.GET("/checklists/:id", getChecklistByID)
	router.PATCH("/checklists/:id", editChecklistByID)
	router.DELETE("/checklists/:id", removeChecklistByID)
	router.POST("/checklists/:id/reset", resetChecklistByID)

	router.GET("/checklist-items", getChecklistItems)
	router.POST("/checklist-items/:id", addChecklistItem)
	router.PATCH("/checklist-items/:id", editChecklistItemByID)
	router.DELETE("/checklist-items/:id", removeChecklistItemByID)

	router.Run("localhost:8080")
}
