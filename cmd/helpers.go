package main

import (
	"strconv"
	"strings"
	"time"
)

func addIdToChecklist(addChecklist *AddChecklist) Checklist {
	checklist := Checklist{
		ID:    generateSlug(addChecklist.Title),
		Title: addChecklist.Title,
	}
	return checklist
}

func generateSlug(title string) string {
	slug := strings.ReplaceAll(strings.ToLower(title), " ", "-")

	for _, checklist := range checklists {
		if checklist.ID == slug {
			slug = slug + strconv.FormatInt(time.Now().UnixNano(), 10)
			break
		}
	}
	return slug
}
