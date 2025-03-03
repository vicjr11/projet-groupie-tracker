package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet-groupie-tracker/models"
	"time"
)



func GetLocation() (models.ListLocation, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "https://rickandmortyapi.com/api/location", nil)
	if reqErr != nil {
		return models.ListLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return models.ListLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.ListLocation{}, res.StatusCode, fmt.Errorf("GetLocation - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data models.ListLocation
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return models.ListLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

//

func LocationPage(page int) (models.ListLocation, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	// Construire l'URL avec le paramètre de page
	url := fmt.Sprintf("https://rickandmortyapi.com/api/location?page=%d", page)
	
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return models.ListLocation{}, http.StatusInternalServerError, fmt.Errorf("LocationPage - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return models.ListLocation{}, http.StatusInternalServerError, fmt.Errorf("LocationPage - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.ListLocation{}, res.StatusCode, fmt.Errorf("LocationPage - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data models.ListLocation
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return models.ListLocation{}, http.StatusInternalServerError, fmt.Errorf("LocationPage - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}