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
		http.HandleFunc("/favoris", controllers.FavorisControllers) // Configure la route pour les favoris
		http.HandleFunc("/favoris/add-character", controllers.AddCharacterToFavorites)
		http.HandleFunc("/favoris/remove-character", controllers.RemoveCharacterFromFavorites)
		http.HandleFunc("/favoris/add-episode", controllers.AddEpisodeToFavorites)
		http.HandleFunc("/favoris/remove-episode", controllers.RemoveEpisodeFromFavorites)
		http.HandleFunc("/favoris/add-location", controllers.AddLocationToFavorites)
		http.HandleFunc("/favoris/remove-location", controllers.RemoveLocationFromFavorites)
		http.HandleFunc("/favoris/check", controllers.CheckFavoriteStatus)
	}


func DetailRoutes() {
//	http.HandleFunc("/details",controllers.CharacterDetailController)  // configure la route pour les favoris
}
func SearchRoutes() {
	http.HandleFunc("/search", controllers.SearchControllers) // Configure la route pour la recherche
}
func utilsRoutes(){
	//http.HandleFunc("/utils",controllers.UtilsControllers)
}

