package controllers


import (
	"net/http"
	"projet-groupie-tracker/templates"
)

func FavorisControllers(w http.ResponseWriter, r *http.Request) {
temp.Temp.ExecuteTemplate(w, "favoris", nil)
}