package controllers

import (
	"net/http"
	"projet-groupie-tracker/models"
	"projet-groupie-tracker/services"
	"projet-groupie-tracker/templates"
	
)

// SearchControllers handles search requests
func SearchControllers(w http.ResponseWriter, r *http.Request) {
	// Get search query from parameters
	query := r.URL.Query().Get("q")

	// If query is empty, display empty search page
	if query == "" {
		data := struct {
			Query   string
			Results []models.SearchResult
			Count   int
		}{
			Query:   "",
			Results: []models.SearchResult{},
			Count:   0,
		}
		temp.Temp.ExecuteTemplate(w, "search", data)
		return
	}

	// Perform search
	results, err := services.SearchAll(query)
	if err != nil {
		http.Error(w, "Error during search: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare data for template
	data := struct {
		Query   string
		Results []models.SearchResult
		Count   int
	}{
		Query:   query,
		Results: results,
		Count:   len(results),
	}

	// Display results using the template system
	temp.Temp.ExecuteTemplate(w, "search", data)
}
