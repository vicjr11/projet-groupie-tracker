package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"projet-groupie-tracker/models"
	"time"
)

// HTTPClient interface for easier testing
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// DefaultHTTPClient returns a standard HTTP client with timeout
func DefaultHTTPClient() HTTPClient {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

// PaginationInfo contient les informations de pagination
type PaginationInfo struct {
	CurrentPage  int
	TotalPages   int
	HasPrevPage  bool
	HasNextPage  bool
	TotalResults int
	PageSize     int
}

// CalculatePagination calcule les informations de pagination
func CalculatePagination(page, totalResults, pageSize int) PaginationInfo {
	totalPages := int(math.Ceil(float64(totalResults) / float64(pageSize)))
	
	// Ajuster la page si elle dépasse les limites
	if page < 1 {
		page = 1
	}
	if page > totalPages && totalPages > 0 {
		page = totalPages
	}
	
	return PaginationInfo{
		CurrentPage:  page,
		TotalPages:   totalPages,
		HasPrevPage:  page > 1,
		HasNextPage:  page < totalPages,
		TotalResults: totalResults,
		PageSize:     pageSize,
	}
}

// GeneratePageNumbers génère une liste de numéros de page à afficher
func GeneratePageNumbers(currentPage, totalPages, maxVisible int) []int {
	if totalPages <= maxVisible {
		// Si le nombre total de pages est inférieur au nombre max visible,
		// afficher toutes les pages
		result := make([]int, totalPages)
		for i := 0; i < totalPages; i++ {
			result[i] = i + 1
		}
		return result
	}
	
	// Sinon, calculer quelles pages afficher
	result := make([]int, 0, maxVisible)
	
	// Toujours afficher la première page
	result = append(result, 1)
	
	// Calculer la plage de pages autour de la page courante
	startPage := currentPage - maxVisible/2
	endPage := currentPage + maxVisible/2
	
	// Ajuster si nécessaire
	if startPage <= 1 {
		startPage = 2
		endPage = min(totalPages-1, maxVisible-1)
	} else if endPage >= totalPages {
		endPage = totalPages - 1
		startPage = max(2, totalPages-maxVisible+2)
	}
	
	// Ajouter des points de suspension si nécessaire
	if startPage > 2 {
		result = append(result, -1) // -1 représente "..."
	}
	
	// Ajouter les pages du milieu
	for i := startPage; i <= endPage; i++ {
		result = append(result, i)
	}
	
	// Ajouter des points de suspension si nécessaire
	if endPage < totalPages-1 {
		result = append(result, -1) // -1 représente "..."
	}
	
	// Toujours afficher la dernière page
	if totalPages > 1 {
		result = append(result, totalPages)
	}
	
	return result
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// PaginatedResponse is a generic interface for responses with pagination
type PaginatedResponse interface {
	HasNextPage() bool
}

// GetAllPages fetches all pages from a paginated API endpoint
func GetAllPages[T any](baseURL string) ([]T, int, error) {
	var allResults []T
	page := 1

	for {
		// Fetch current page
		pageData, statusCode, err := GetPage[T](baseURL, page)
		if err != nil {
			return nil, statusCode, err
		}

		// Extract and append results
		results, hasNext := extractResults(pageData)
		allResults = append(allResults, results...)

		// Check if there are more pages
		if !hasNext {
			break
		}

		page++
	}

	return allResults, http.StatusOK, nil
}

// GetPage fetches a specific page from an API endpoint
func GetPage[T any](url string, page int) (T, int, error) {
	var result T
	client := DefaultHTTPClient()

	// Create request for specific page
	req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?page=%d", url, page), nil)
	if reqErr != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("GetPage - Error preparing request: %s", reqErr)
	}

	// Send request
	res, resErr := client.Do(req)
	if resErr != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("GetPage - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	// Check response status
	if res.StatusCode != http.StatusOK {
		return result, res.StatusCode, fmt.Errorf("GetPage - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	// Decode response
	decodeErr := json.NewDecoder(res.Body).Decode(&result)
	if decodeErr != nil {
		return result, http.StatusInternalServerError, fmt.Errorf("GetPage - Error decoding data: %s", decodeErr.Error())
	}

	return result, res.StatusCode, nil
}

// extractResults extracts the results from a paginated response and checks if there's a next page
// This function is generic and works with different response types
func extractResults[T any](response T) ([]T, bool) {
	// Implement for specific response types
	switch v := any(response).(type) {
	case models.ListCharacters:
		var results []models.Character
		for _, item := range v.Results {
			results = append(results, item)
		}
		return any(results).([]T), v.Info.Next != ""
	case models.ListEpisode:
		var results []models.Episode
		for _, item := range v.Results {
			results = append(results, item)
		}
		return any(results).([]T), v.Info.Next != ""
	case models.ListLocation:
		var results []models.Location
		for _, item := range v.Results {
			results = append(results, item)
		}
		return any(results).([]T), v.Info.Next != ""
	case ResultsCharacter:
		return any(v.Results).([]T), v.Info.Next != ""
	case ResultsEpisode:
		return any(v.Results).([]T), v.Info.Next != ""
	case ResultsLocation:
		return any(v.Results).([]T), v.Info.Next != ""
	default:
		// Return empty result and no next page for unknown types
		var empty []T
		return empty, false
	}
}

// ResultsEpisode structure pour les résultats des épisodes avec pagination
type ResultsEpisode struct {
	Info struct {
		Count int    `json:"count"`
		Pages int    `json:"pages"`
		Next  string `json:"next"`
		Prev  string `json:"prev"`
	} `json:"info"`
	Results []models.ItemEpisode `json:"results"`
}

func GetAllepisode() ([]models.ItemEpisode, int, error) {
	var ListEpisode []models.ItemEpisode

	_client := http.Client{
		Timeout: 5 * time.Second,
	}
	page := 1
	for {
		req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/episode?page=%d", page), nil)
		if reqErr != nil {
			return []models.ItemEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error preparing request: %s", reqErr)
		}

		res, resErr := _client.Do(req)
		if resErr != nil {
			return []models.ItemEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpidose - Error sending request: %s", resErr)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return []models.ItemEpisode{}, res.StatusCode, fmt.Errorf("GetEpisode - Error in response code: %d, message: %s", res.StatusCode, res.Status)
		}

		var data ResultsEpisode
		decodeErr := json.NewDecoder(res.Body).Decode(&data)
		if decodeErr != nil {
			return []models.ItemEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error decoding data: %s", decodeErr.Error())
		}
		ListEpisode = append(ListEpisode, data.Results...)
		page++
		if data.Info.Next == "" {
			break
		}
	}

	return ListEpisode, http.StatusOK, nil
}

func GetepisodePage(page int) (models.ListEpisode, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/episode?page=%d", page), nil)
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

// ResultsLocation structure pour les résultats des lieux avec pagination
type ResultsLocation struct {
	Info struct {
		Count int    `json:"count"`
		Pages int    `json:"pages"`
		Next  string `json:"next"`
		Prev  string `json:"prev"`
	} `json:"info"`
	Results []models.ItemLocation `json:"results"`
}

func GetAlllocation() ([]models.ItemLocation, int, error) {
	var ListLocation []models.ItemLocation

	_client := http.Client{
		Timeout: 5 * time.Second,
	}
	page := 1
	for {
		req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/location?page=%d", page), nil)
		if reqErr != nil {
			return []models.ItemLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error preparing request: %s", reqErr)
		}

		res, resErr := _client.Do(req)
		if resErr != nil {
			return []models.ItemLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error sending request: %s", resErr)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return []models.ItemLocation{}, res.StatusCode, fmt.Errorf("GetLocation - Error in response code: %d, message: %s", res.StatusCode, res.Status)
		}

		var data ResultsLocation
		decodeErr := json.NewDecoder(res.Body).Decode(&data)
		if decodeErr != nil {
			return []models.ItemLocation{}, http.StatusInternalServerError, fmt.Errorf("GetLocation - Error decoding data: %s", decodeErr.Error())
		}
		ListLocation = append(ListLocation, data.Results...)
		page++
		if data.Info.Next == "" {
			break
		}
	}

	return ListLocation, http.StatusOK, nil
}

// GetLocationPage returns locations for a specific page
func GetLocationPage(page int) (models.ListLocation, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/location?page=%d", page), nil)
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

// ItemCharacter structure pour les informations d'un personnage
type ItemCharacter struct {
    ID       int      `json:"id"`
    Name     string   `json:"name"`
    Species  string   `json:"species"`
    Status   string   `json:"status"`
    Gender   string   `json:"gender"`
    Origin   struct {
        Name string `json:"name"`
        Url  string `json:"url"`
    } `json:"origin"`
    Location struct {
        Name string `json:"name"`
        Url  string `json:"url"`
    } `json:"location"`
    Image    string   `json:"image"`
    Episode  []string `json:"episode"`
    URL      string   `json:"url"`
}

// ResultsCharacter structure pour les résultats des personnages avec pagination
type ResultsCharacter struct {
	Info struct {
		Count int    `json:"count"`
		Pages int    `json:"pages"`
		Next  string `json:"next"`
		Prev  string `json:"prev"`
	} `json:"info"`
	Results []ItemCharacter `json:"results"`
}

func GetAllcharacter() ([]ItemCharacter, int, error) {
	var listCharacters []ItemCharacter

	_client := http.Client{
		Timeout: 5 * time.Second,
	}
	page := 1
	for {
		req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", page), nil)
		if reqErr != nil {
			return []ItemCharacter{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error preparing request: %s", reqErr)
		}

		res, resErr := _client.Do(req)
		if resErr != nil {
			return []ItemCharacter{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error sending request: %s", resErr)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return []ItemCharacter{}, res.StatusCode, fmt.Errorf("GetCharacter - Error in response code: %d, message: %s", res.StatusCode, res.Status)
		}

		var data ResultsCharacter
		decodeErr := json.NewDecoder(res.Body).Decode(&data)
		if decodeErr != nil {
			return []ItemCharacter{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error decoding data: %s", decodeErr.Error())
		}
		listCharacters = append(listCharacters, data.Results...)
		page++
		if data.Info.Next == "" {
			break
		}
	}

	return listCharacters, http.StatusOK, nil
}

// GetcharacterPage returns characters for a specific page
func GetcharacterPage(page int) (ResultsCharacter, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", page), nil)
	if reqErr != nil {
		return ResultsCharacter{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return ResultsCharacter{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ResultsCharacter{}, res.StatusCode, fmt.Errorf("GetCharacter - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data ResultsCharacter
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return ResultsCharacter{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}