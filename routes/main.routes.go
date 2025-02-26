package routes

import (
	"fmt"
	"net/http"
	"projet-groupie-tracker/controllers"
)

func MainServe() {
	http.HandleFunc("/", controllers.HomeControllers)

	ProjetsRoutes()
	CharacterRoutes()
	EpisodeRoutes()
	LocationRoutes()
	FavorisRoutes()
	DetailRoutes()
	SearchRoutes() 
	fmt.Println("Serveur lancé : http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
	//test

}
