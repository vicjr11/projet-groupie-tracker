package controllers

import (
	"encoding/json"
	"net/http"
	"projet-groupie-tracker/services"
	"projet-groupie-tracker/templates"
	"strconv"
)

// FavorisControllers handles the favorites page rendering
func FavorisControllers(w http.ResponseWriter, r *http.Request) {
	favoritesManager := services.GetFavoritesManager()
	favorites := favoritesManager.GetAllFavorites()
	
	temp.Temp.ExecuteTemplate(w, "favoris", favorites)
}

// AddCharacterToFavorites handles adding a character to favorites
func AddCharacterToFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	characterIdStr := r.FormValue("character_id")
	if characterIdStr == "" {
		http.Error(w, "Character ID is required", http.StatusBadRequest)
		return
	}

	characterId, err := strconv.Atoi(characterIdStr)
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	favoritesManager := services.GetFavoritesManager()
	err = favoritesManager.AddCharacterToFavorites(characterId)
	
	response := map[string]interface{}{
		"success": err == nil,
	}
	
	if err != nil {
		response["error"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RemoveCharacterFromFavorites handles removing a character from favorites
func RemoveCharacterFromFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	characterIdStr := r.FormValue("character_id")
	if characterIdStr == "" {
		http.Error(w, "Character ID is required", http.StatusBadRequest)
		return
	}

	characterId, err := strconv.Atoi(characterIdStr)
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	favoritesManager := services.GetFavoritesManager()
	err = favoritesManager.RemoveCharacterFromFavorites(characterId)
	
	response := map[string]interface{}{
		"success": err == nil,
	}
	
	if err != nil {
		response["error"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AddEpisodeToFavorites handles adding an episode to favorites
func AddEpisodeToFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	episodeURL := r.FormValue("episode_url")
	if episodeURL == "" {
		http.Error(w, "Episode URL is required", http.StatusBadRequest)
		return
	}

	favoritesManager := services.GetFavoritesManager()
	err := favoritesManager.AddEpisodeToFavorites(episodeURL)
	
	response := map[string]interface{}{
		"success": err == nil,
	}
	
	if err != nil {
		response["error"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RemoveEpisodeFromFavorites handles removing an episode from favorites
func RemoveEpisodeFromFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	episodeIdStr := r.FormValue("episode_id")
	if episodeIdStr == "" {
		http.Error(w, "Episode ID is required", http.StatusBadRequest)
		return
	}

	episodeId, err := strconv.Atoi(episodeIdStr)
	if err != nil {
		http.Error(w, "Invalid episode ID", http.StatusBadRequest)
		return
	}

	favoritesManager := services.GetFavoritesManager()
	err = favoritesManager.RemoveEpisodeFromFavorites(episodeId)
	
	response := map[string]interface{}{
		"success": err == nil,
	}
	
	if err != nil {
		response["error"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AddLocationToFavorites handles adding a location to favorites
func AddLocationToFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	locationURL := r.FormValue("location_url")
	if locationURL == "" {
		http.Error(w, "Location URL is required", http.StatusBadRequest)
		return
	}

	favoritesManager := services.GetFavoritesManager()
	err := favoritesManager.AddLocationToFavorites(locationURL)
	
	response := map[string]interface{}{
		"success": err == nil,
	}
	
	if err != nil {
		response["error"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RemoveLocationFromFavorites handles removing a location from favorites
func RemoveLocationFromFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	locationIdStr := r.FormValue("location_id")
	if locationIdStr == "" {
		http.Error(w, "Location ID is required", http.StatusBadRequest)
		return
	}

	locationId, err := strconv.Atoi(locationIdStr)
	if err != nil {
		http.Error(w, "Invalid location ID", http.StatusBadRequest)
		return
	}

	favoritesManager := services.GetFavoritesManager()
	err = favoritesManager.RemoveLocationFromFavorites(locationId)
	
	response := map[string]interface{}{
		"success": err == nil,
	}
	
	if err != nil {
		response["error"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CheckFavoriteStatus checks if an item is in favorites
func CheckFavoriteStatus(w http.ResponseWriter, r *http.Request) {
	itemType := r.URL.Query().Get("type")
	idStr := r.URL.Query().Get("id")
	
	if itemType == "" || idStr == "" {
		http.Error(w, "Item type and ID are required", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	
	favoritesManager := services.GetFavoritesManager()
	var isFavorite bool
	
	switch itemType {
	case "character":
		isFavorite = favoritesManager.IsCharacterFavorite(id)
	default:
		http.Error(w, "Invalid item type", http.StatusBadRequest)
		return
	}
	
	response := map[string]interface{}{
		"is_favorite": isFavorite,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
