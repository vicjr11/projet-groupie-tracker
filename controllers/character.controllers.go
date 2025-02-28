package controllers

import (
	"net/http"
	"projet-groupie-tracker/models"
	"projet-groupie-tracker/services"
	temp "projet-groupie-tracker/templates"
	"strconv"
)

func CharactersControllers(w http.ResponseWriter, r *http.Request) {
	// Get page from query parameter, default to 1
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
	}

	// Get characters for specified page
	listCharacters, statusCode, errors := services.CharacterPage(page)
	if errors != nil {
		http.Error(w, errors.Error(), statusCode)
		return
	}

	// Creation de la pagination data
	paginationData := struct {
		Data      models.ListCharacters
		Page      int
		HasPrev   bool
		HasNext   bool
		TotalPage int
	}{
		Data:      listCharacters,
		Page:      page,
		HasPrev:   page > 1,
		HasNext:   page < listCharacters.Info.Pages,
		TotalPage: listCharacters.Info.Pages,
	}

	temp.Temp.ExecuteTemplate(w, "character", paginationData)
}
