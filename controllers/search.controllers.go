package controllers

import (
	"html/template"
	"net/http"
	"projet-groupie-tracker/services"
)

// SearchControllers gère les requêtes de recherche
func SearchControllers(w http.ResponseWriter, r *http.Request) {
	// Récupérer la requête de recherche depuis les paramètres
	query := r.URL.Query().Get("q")
	
	// Si la requête est vide, afficher simplement la page de recherche
	if query == "" {
		tmpl := template.Must(template.ParseFiles("templates/search.html"))
		tmpl.Execute(w, nil)
		return
	}
	
	// Effectuer la recherche
	results, err := services.SearchAll(query)
	if err != nil {
		http.Error(w, "Erreur lors de la recherche: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Préparer les données pour le template
	data := struct {
		Query   string
		Results []services.SearchResult
		Count   int
	}{
		Query:   query,
		Results: results,
		Count:   len(results),
	}
	
	// Afficher les résultats
	tmpl := template.Must(template.ParseFiles("templates/search.html"))
	tmpl.Execute(w, data)
}