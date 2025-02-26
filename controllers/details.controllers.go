package controllers

import (
	"net/http"
	"projet-groupie-tracker/services"
	temp "projet-groupie-tracker/templates"
	"strconv"
)

type CharacterDetailData struct {
	Character    services.CharacterDetail
	Episodes     []services.Episode
	OriginInfo   map[string]interface{}
	LocationInfo map[string]interface{}
}

func CharacterDetailController(w http.ResponseWriter, r *http.Request) {
	// Get character ID from URL parameter
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Character ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	// Get character details
	character, statusCode, detailErr := services.CharacterDetail(id)
	if detailErr != nil {
		http.Error(w, detailErr.Error(), statusCode)
		return
	}

	// Get episodes where character appears
	episodes, epStatusCode, epErr := services.GetCharacterEpisodes(character.Episode)
	if epErr != nil {
		http.Error(w, epErr.Error(), epStatusCode)
		return
	}

	// Get origin location details
	originInfo, originStatusCode, originErr := services.GetLocationDetails(character.Origin.Url)
	if originErr != nil {
		http.Error(w, originErr.Error(), originStatusCode)
		return
	}

	// Get current location details
	locationInfo, locStatusCode, locErr := services.GetLocationDetails(character.Location.Url)
	if locErr != nil {
		http.Error(w, locErr.Error(), locStatusCode)
		return
	}

	// Prepare data for template
	detailData := CharacterDetailData{
		Character:    character,
		Episodes:     episodes,
		OriginInfo:   originInfo,
		LocationInfo: locationInfo,
	}

	// Render the template
	temp.Temp.ExecuteTemplate(w, "characterDetail", detailData)
}