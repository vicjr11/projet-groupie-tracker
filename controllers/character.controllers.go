package controllers

import (
	"net/http"
	"projet-groupie-tracker/services"
	temp "projet-groupie-tracker/templates"
)

func CharactersControllers(w http.ResponseWriter, r *http.Request) {
	listCharacters, statusCode, errors := services.Getcharacter()
	if errors != nil {
		http.Error(w, errors.Error(), statusCode)
		return
	}
	temp.Temp.ExecuteTemplate(w, "character", listCharacters)
}
