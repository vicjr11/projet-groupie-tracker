package controllers

import (
	"net/http"
	"projet-groupie-tracker/services"
	"projet-groupie-tracker/templates"
)

func HomeControllers(w http.ResponseWriter, r *http.Request) {
	services.GetAllcharacter()
	temp.Temp.ExecuteTemplate(w, "home", nil)
}