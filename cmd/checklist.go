package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Checklist struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type AddChecklist struct {
	Title string `json:"title"`
}

func getChecklists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, checklists)
}

func addChecklist(c *gin.Context) {
	var checklistToAdd AddChecklist

	err := c.BindJSON(&checklistToAdd)
	if err != nil {
		return
	}

	newChecklist := addIdToChecklist(&checklistToAdd)
	checklists = append(checklists, newChecklist)
	c.IndentedJSON(http.StatusCreated, newChecklist)
}

func getChecklistByID(c *gin.Context) {
	id := c.Param("id")

	for _, checklist := range checklists {
		if checklist.ID == id {
			c.IndentedJSON(http.StatusOK, checklist)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "checklist not found"})
}

func editChecklistByID(c *gin.Context) {
	id := c.Param("id")
	var editChecklist AddChecklist

	err := c.BindJSON(&editChecklist)
	if err != nil {
		return
	}

	for i, checklist := range checklists {
		if checklist.ID == id {
			checklists[i].Title = editChecklist.Title
			c.IndentedJSON(http.StatusOK, editChecklist)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "checklist does not exist"})
}

func removeChecklistByID(c *gin.Context) {
	id := c.Param("id")

	for i, checklist := range checklists {
		if checklist.ID == id {
			checklists = append(checklists[:i], checklists[i+1:]...)
			return
		}
	}
}
