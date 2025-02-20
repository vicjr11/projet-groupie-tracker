package controllers

import (
	"net/http"
	"projet-groupie-tracker/services"
	temp "projet-groupie-tracker/templates"
)

func HomeControllers(w http.ResponseWriter, r *http.Request) {
	services.GetAllepisode()
	temp.Temp.ExecuteTemplate(w, "home", nil)
}
