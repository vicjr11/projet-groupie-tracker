package routes

import (
	"net/http"
	"projet-groupie-tracker/controllers"
)

func ProjetsRoutes() {
	http.HandleFunc("/index", controllers.HomeControllers) // Configure les routes pour la section projets
	//http.HandleFunc("/projets", nil)                       // Inclut la page d'accueil et la page des projets
}

func CharacterRoutes() {
	http.HandleFunc("/character", controllers.CharactersControllers) // Configure la route pour le character
}
func EpisodeRoutes() {
	http.HandleFunc("/episode", controllers.EpisodeControllers) // Configure la route pour les episodes
}

func LocationRoutes() {
	http.HandleFunc("/location", controllers.LocationControllers) // Configure la route pour les locations
}

func FavorisRoutes() {
	http.HandleFunc("/favoris", controllers.FavorisControllers) // configure la route pour les favoris
}

func DetailRoutes() {
	http.HandleFunc("/details",controllers.DetailsControllers)  // configure la route pour les favoris
}
func SearchRoutes() {
	http.HandleFunc("/search", controllers.SearchControllers) // Configure la route pour la recherche
}