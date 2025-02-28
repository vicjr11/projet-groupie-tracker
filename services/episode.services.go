package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet-groupie-tracker/models"
	"time"
)

func Getepisode() (models.ListEpisode, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "https://rickandmortyapi.com/api/episode", nil)
	if reqErr != nil {
		return models.ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return models.ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.ListEpisode{}, res.StatusCode, fmt.Errorf("GetEpisode - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data models.ListEpisode
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return models.ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

// GetepisodePage récupère une page spécifique d'épisodes
func GetepisodePage(page int) (models.ListEpisode, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	// Construire l'URL avec le paramètre de page
	url := fmt.Sprintf("https://rickandmortyapi.com/api/episode?page=%d", page)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return models.ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetepisodePage - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return models.ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetepisodePage - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.ListEpisode{}, res.StatusCode, fmt.Errorf("GetepisodePage - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data models.ListEpisode
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return models.ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetepisodePage - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}
