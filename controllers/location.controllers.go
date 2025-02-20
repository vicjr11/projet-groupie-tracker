package controllers

import (
	"net/http"
	"projet-groupie-tracker/services"
	"projet-groupie-tracker/templates"
)

func LocationControllers(w http.ResponseWriter, r *http.Request) {
	ListLocation, statusCode, errors := services.GetLocation()
	if errors != nil {
		http.Error(w, errors.Error(), statusCode)
		return
	}
	temp.Temp.ExecuteTemplate(w, "location", ListLocation)
}
  