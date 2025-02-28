package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"projet-groupie-tracker/models"
	"strings"
)

// FilterParams represents generic filter parameters
type FilterParams map[string]string


// ToFilterParams converts a character filter to generic filter parameters
func (f models.FilterCharacter) ToFilterParams() FilterParams {
	params := make(FilterParams)
	if f.Name != "" {
		params["name"] = f.Name
	}
	if f.Status != "" {
		params["status"] = f.Status
	}
	if f.Species != "" {
		params["species"] = f.Species
	}
	if f.Type != "" {
		params["type"] = f.Type
	}
	if f.Gender != "" {
		params["gender"] = f.Gender
	}
	return params
}

// ToFilterParams converts an episode filter to generic filter parameters
func (f models.FilterEpisode) ToFilterParams() FilterParams {
	params := make(FilterParams)
	if f.Name != "" {
		params["name"] = f.Name
	}
	if f.Episode != "" {
		params["episode"] = f.Episode
	}
	return params
}

// ToFilterParams converts a location filter to generic filter parameters
func (f models.FilterLocation) ToFilterParams() FilterParams {
	params := make(FilterParams)
	if f.Name != "" {
		params["name"] = f.Name
	}
	if f.Type != "" {
		params["type"] = f.Type
	}
	if f.Dimension != "" {
		params["dimension"] = f.Dimension
	}
	return params
}

// FilterResource performs a filtered request on any Rick and Morty API resource
func FilterResource(baseURL string, params FilterParams, client HTTPClient, result interface{}) (int, error) {
	// Build URL with query parameters
	queryURL, err := buildFilterURL(baseURL, params)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("FilterResource - Error building filter URL: %s", err)
	}

	// Create request
	req, reqErr := http.NewRequest(http.MethodGet, queryURL, nil)
	if reqErr != nil {
		return http.StatusInternalServerError, fmt.Errorf("FilterResource - Error preparing request: %s", reqErr)
	}

	// Send request
	res, resErr := client.Do(req)
	if resErr != nil {
		return http.StatusInternalServerError, fmt.Errorf("FilterResource - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	// Check for errors
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, fmt.Errorf("FilterResource - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	// Decode response
	decodeErr := json.NewDecoder(res.Body).Decode(result)
	if decodeErr != nil {
		return http.StatusInternalServerError, fmt.Errorf("FilterResource - Error decoding data: %s", decodeErr.Error())
	}

	return res.StatusCode, nil
}

// buildFilterURL constructs a URL with query parameters
func buildFilterURL(baseURL string, params FilterParams) (string, error) {
	if len(params) == 0 {
		return baseURL, nil
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	query := parsedURL.Query()
	for key, value := range params {
		query.Add(key, value)
	}

	parsedURL.RawQuery = query.Encode()
	return parsedURL.String(), nil
}

// FilterCharacters filters characters based on provided parameters
func FilterCharacters(filter models.FilterCharacter) (models.ListCharacters, int, error) {
	var result models.ListCharacters
	client := DefaultHTTPClient()
	params := filter.ToFilterParams()

	statusCode, err := FilterResource("https://rickandmortyapi.com/api/character", params, client, &result)
	if err != nil {
		return models.ListCharacters{}, statusCode, err
	}

	return result, statusCode, nil
}

// FilterEpisodes filters episodes based on provided parameters
func FilterEpisodes(filter models.FilterEpisode) (models.ListEpisode, int, error) {
	var result models.ListEpisode
	client := DefaultHTTPClient()
	params := filter.ToFilterParams()

	statusCode, err := FilterResource("https://rickandmortyapi.com/api/episode", params, client, &result)
	if err != nil {
		return models.ListEpisode{}, statusCode, err
	}

	return result, statusCode, nil
}

// FilterLocations filters locations based on provided parameters
func FilterLocations(filter models.FilterLocation) (models.ListLocation, int, error) {
	var result models.ListLocation
	client := DefaultHTTPClient()
	params := filter.ToFilterParams()

	statusCode, err := FilterResource("https://rickandmortyapi.com/api/location", params, client, &result)
	if err != nil {
		return models.ListLocation{}, statusCode, err
	}

	return result, statusCode, nil
}

// MultiFilter handles filters with multiple comma-separated values
// The Rick and Morty API allows filtering by multiple IDs (e.g., /character/1,2,3)
func MultiFilter[T any](baseURL string, ids []int, client HTTPClient, result *[]T) (int, error) {
	if len(ids) == 0 {
		return http.StatusBadRequest, fmt.Errorf("MultiFilter - No IDs provided")
	}

	// Convert IDs to comma-separated string
	idStrings := make([]string, len(ids))
	for i, id := range ids {
		idStrings[i] = fmt.Sprintf("%d", id)
	}
	idString := strings.Join(idStrings, ",")

	// Create URL
	filterURL := fmt.Sprintf("%s/%s", baseURL, idString)

	// Create request
	req, reqErr := http.NewRequest(http.MethodGet, filterURL, nil)
	if reqErr != nil {
		return http.StatusInternalServerError, fmt.Errorf("MultiFilter - Error preparing request: %s", reqErr)
	}

	// Send request
	res, resErr := client.Do(req)
	if resErr != nil {
		return http.StatusInternalServerError, fmt.Errorf("MultiFilter - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	// Check for errors
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, fmt.Errorf("MultiFilter - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	// Handle both single item and array responses
	// When requesting a single ID, the API returns an object, not an array
	if len(ids) == 1 {
		var singleItem T
		decodeErr := json.NewDecoder(res.Body).Decode(&singleItem)
		if decodeErr != nil {
			return http.StatusInternalServerError, fmt.Errorf("MultiFilter - Error decoding data: %s", decodeErr.Error())
		}
		*result = []T{singleItem}
	} else {
		// Multiple IDs, decode as array
		decodeErr := json.NewDecoder(res.Body).Decode(result)
		if decodeErr != nil {
			return http.StatusInternalServerError, fmt.Errorf("MultiFilter - Error decoding data: %s", decodeErr.Error())
		}
	}

	return res.StatusCode, nil
}

// GetCharactersByID fetches multiple characters by their IDs
func GetCharactersByID(ids []int) ([]models.Character, int, error) {
	var results []models.Character
	client := DefaultHTTPClient()

	statusCode, err := MultiFilter[models.Character]("https://rickandmortyapi.com/api/character", ids, client, &results)
	if err != nil {
		return nil, statusCode, err
	}

	return results, statusCode, nil
}

// GetEpisodesByID fetches multiple episodes by their IDs
func GetEpisodesByID(ids []int) ([]models.Episode, int, error) {
	var results []models.Episode
	client := DefaultHTTPClient()

	statusCode, err := MultiFilter[models.Episode]("https://rickandmortyapi.com/api/episode", ids, client, &results)
	if err != nil {
		return nil, statusCode, err
	}

	return results, statusCode, nil
}

// GetLocationsByID fetches multiple locations by their IDs
func GetLocationsByID(ids []int) ([]models.Location, int, error) {
	var results []models.Location
	client := DefaultHTTPClient()

	statusCode, err := MultiFilter[models.Location]("https://rickandmortyapi.com/api/location", ids, client, &results)
	if err != nil {
		return nil, statusCode, err
	}

	return results, statusCode, nil
}
