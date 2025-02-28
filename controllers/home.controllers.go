package controllers

import (
	"net/http"
	temp "projet-groupie-tracker/templates"
	"projet-groupie-tracker/utils"
)

func HomeControllers(w http.ResponseWriter, r *http.Request) {
	utils.GetAllepisode()
	temp.Temp.ExecuteTemplate(w, "home", nil)
}
