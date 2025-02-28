package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Location struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	Residents string `json:"resident"`
	Url       string  `json:"url"`

}

type ListLocation struct {
	Info struct {
		Count int
		Pages int
		Next  string
	}
	Results []Location
}

func GetLocation() (ListLocation, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "https://rickandmortyapi.com/api/location", nil)
	if reqErr != nil {
		return ListLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return ListLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ListLocation{}, res.StatusCode, fmt.Errorf("GetLocation - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data ListLocation
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return ListLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

//

func LocationPage(page int) (ListLocation, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	// Construire l'URL avec le paramètre de page
	url := fmt.Sprintf("https://rickandmortyapi.com/api/location?page=%d", page)
	
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return ListLocation{}, http.StatusInternalServerError, fmt.Errorf("LocationPage - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return ListLocation{}, http.StatusInternalServerError, fmt.Errorf("LocationPage - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ListLocation{}, res.StatusCode, fmt.Errorf("LocationPage - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data ListLocation
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return ListLocation{}, http.StatusInternalServerError, fmt.Errorf("LocationPage - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}