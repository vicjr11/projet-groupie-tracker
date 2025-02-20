package controllers

import (
	"net/http"
	"projet-groupie-tracker/services"
	"projet-groupie-tracker/templates"
)

func EpisodeControllers(w http.ResponseWriter, r *http.Request) {
	ListEpisode, statusCode, errors := services.Getepisode()
	if errors != nil {
		http.Error(w, errors.Error(), statusCode)
		return
	}
	temp.Temp.ExecuteTemplate(w, "episode", ListEpisode)
}
