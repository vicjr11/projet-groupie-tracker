package controllers

import (
		"net/http"
		"projet-groupie-tracker/templates"
)

func HomeControllers(w http.ResponseWriter, r *http.Request) {
	temp.Temp.ExecuteTemplate(w, "home", nil)
}