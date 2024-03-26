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

	router.Use(cors.Default())
	router.GET("/checklists", getChecklists)
	router.POST("/checklists", addChecklist)

	router.GET("/checklists/:id", getChecklistByID)
	router.POST("/checklists/:id", editChecklistByID)
	router.DELETE("/checklists/:id", removeChecklistByID)
	router.POST("/checklists/:id/reset", resetChecklistByID)

	router.GET("/checklist-items/:id", getItemsByChecklistID)
	router.POST("/checklist-items", addChecklistItem)
	router.POST("/checklist-items/:id", editChecklistItemByID)
	router.DELETE("/checklist-items/:id", removeChecklistItemByID)

	router.Run("localhost:8080")
}
