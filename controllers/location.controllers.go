package controllers

import (
	"net/http"
	"projet-groupie-tracker/models"
	"projet-groupie-tracker/services"
	temp "projet-groupie-tracker/templates"
	"strconv"
)

func LocationControllers(w http.ResponseWriter, r *http.Request) {
	// LocationControllers with pagination
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
	listLocation, statusCode, errors := services.LocationPage(page)
	if errors != nil {
		http.Error(w, errors.Error(), statusCode)
		return
	}

	paginationData := struct {
		Data        models.ListLocation
		Page        int
		HasPrev     bool
		HasNext     bool
		PrevPage    int
		NextPage    int
		TotalPage   int
		PageNumbers []int // Ajout de cette ligne
	}{
		Data:        listLocation,
		Page:        page,
		HasPrev:     page > 1,
		HasNext:     page < listLocation.Info.Pages,
		PrevPage:    page - 1,
		NextPage:    page + 1,
		TotalPage:   listLocation.Info.Pages,
	}
	

	temp.Temp.ExecuteTemplate(w, "location", paginationData)
}


