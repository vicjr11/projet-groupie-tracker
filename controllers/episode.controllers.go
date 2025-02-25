package controllers

import (
	"net/http"
	"projet-groupie-tracker/services"
	"projet-groupie-tracker/templates"
	"strconv"
)

func EpisodeControllers(w http.ResponseWriter, r *http.Request) {
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
	listEpisode, statusCode, errors := services.GetepisodePage(page)
	if errors != nil {
		http.Error(w, errors.Error(), statusCode)
		return
	}
	
	paginationData := struct {
		Data      services.ListEpisode
		Page      int
		HasPrev   bool
		HasNext   bool
		TotalPage int
	}{
		Data:      listEpisode,
		Page:      page,
		HasPrev:   page > 1,
		HasNext:   page < listEpisode.Info.Pages,
		TotalPage: listEpisode.Info.Pages,
	}
	
	temp.Temp.ExecuteTemplate(w, "episode" , paginationData)
}