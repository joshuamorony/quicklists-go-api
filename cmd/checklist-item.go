package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ChecklistItem struct {
	ID          string `json:"id"`
	ChecklistID string `json:"checklistId"`
	Title       string `json:"title"`
	Checked     bool   `json:"checked"`
}

type AddChecklistItem struct {
	Title string `json:"title"`
}

type EditChecklistItem struct {
	Title   *string `json:"title,omitempty"`
	Checked *bool   `json:"checked,omitempty"`
}

func getItemsByChecklistID(c *gin.Context) {
	id := c.Param("id")

	var matchingItems []ChecklistItem

	for _, checklistItem := range checklistItems {
		if checklistItem.ChecklistID == id {
			matchingItems = append(matchingItems, checklistItem)
		}
	}

	c.IndentedJSON(http.StatusOK, matchingItems)
}

func addChecklistItem(c *gin.Context) {
	id := c.Param("id")
	var checklistItemToAdd AddChecklistItem

	err := c.BindJSON(&checklistItemToAdd)
	if err != nil {
		return
	}

	newChecklistItem := ChecklistItem{
		ID:          time.Now().String(),
		ChecklistID: id,
		Title:       checklistItemToAdd.Title,
		Checked:     false,
	}
	checklistItems = append(checklistItems, newChecklistItem)
	c.IndentedJSON(http.StatusCreated, newChecklistItem)
}

func editChecklistItemByID(c *gin.Context) {
	id := c.Param("id")
	var editChecklistItem EditChecklistItem

	err := c.BindJSON(&editChecklistItem)
	if err != nil {
		return
	}

	for _, checklistItem := range checklistItems {
		if checklistItem.ID == id {
			if editChecklistItem.Title != nil {
				checklistItem.Title = *editChecklistItem.Title
			}
			if editChecklistItem.Checked != nil {
				checklistItem.Checked = *editChecklistItem.Checked
			}
			c.IndentedJSON(http.StatusOK, checklistItem)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "checklist not found"})
}

func removeChecklistItemByID(c *gin.Context) {
	id := c.Param("id")

	for i, checklistItem := range checklistItems {
		if checklistItem.ID == id {
			checklistItems = append(checklistItems[:i], checklistItems[i+1:]...)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "checklist not found"})
}

func resetChecklistByID(c *gin.Context) {
	id := c.Param("id")

	for i, checklistItem := range checklistItems {
		if checklistItem.ChecklistID == id {
			checklistItems[i].Checked = false
		}
	}
}
