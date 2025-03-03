package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet-groupie-tracker/models"
	"time"
	"strings"
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



// LocationFilterOptions contains all possible filter criteria for locations
type LocationFilterOptions struct {
	Type      string
	Dimension string
	SortBy    string
	SortOrder string
}

// FilterLocations filters a list of locations based on provided criteria
func FilterLocations(locations []models.ItemLocation, options LocationFilterOptions) []models.ItemLocation {
	if options.Type == "" && options.Dimension == "" {
		// No filters applied, but we may still need to sort
		if options.SortBy != "" {
			SortLocations(locations, options.SortBy, options.SortOrder)
		}
		return locations
	}
	
	filtered := []models.ItemLocation{}
	
	for _, location := range locations {
		// Type filter
		if options.Type != "" && !strings.EqualFold(location.Type, options.Type) {
			continue
		}
		
		// Dimension filter
		if options.Dimension != "" && !strings.EqualFold(location.Dimension, options.Dimension) {
			continue
		}
		
		// Location passed all filters
		filtered = append(filtered, location)
	}
	
	// Apply sorting if specified
	if options.SortBy != "" {
		SortLocations(filtered, options.SortBy, options.SortOrder)
	}
	
	return filtered
}

// SortLocations sorts a list of locations based on the specified field and order
func SortLocations(locations []models.ItemLocation, sortBy, sortOrder string) {
	switch sortBy {
	case "name":
		if sortOrder == "desc" {
			// Sort by name descending
			for i := 0; i < len(locations)-1; i++ {
				for j := i + 1; j < len(locations); j++ {
					if locations[i].Name < locations[j].Name {
						locations[i], locations[j] = locations[j], locations[i]
					}
				}
			}
		} else {
			// Sort by name ascending (default)
			for i := 0; i < len(locations)-1; i++ {
				for j := i + 1; j < len(locations); j++ {
					if locations[i].Name > locations[j].Name {
						locations[i], locations[j] = locations[j], locations[i]
					}
				}
			}
		}
	case "type":
		if sortOrder == "desc" {
			// Sort by type descending
			for i := 0; i < len(locations)-1; i++ {
				for j := i + 1; j < len(locations); j++ {
					if locations[i].Type < locations[j].Type {
						locations[i], locations[j] = locations[j], locations[i]
					}
				}
			}
		} else {
			// Sort by type ascending (default)
			for i := 0; i < len(locations)-1; i++ {
				for j := i + 1; j < len(locations); j++ {
					if locations[i].Type > locations[j].Type {
						locations[i], locations[j] = locations[j], locations[i]
					}
				}
			}
		}
	case "dimension":
		if sortOrder == "desc" {
			// Sort by dimension descending
			for i := 0; i < len(locations)-1; i++ {
				for j := i + 1; j < len(locations); j++ {
					if locations[i].Dimension < locations[j].Dimension {
						locations[i], locations[j] = locations[j], locations[i]
					}
				}
			}
		} else {
			// Sort by dimension ascending (default)
			for i := 0; i < len(locations)-1; i++ {
				for j := i + 1; j < len(locations); j++ {
					if locations[i].Dimension > locations[j].Dimension {
						locations[i], locations[j] = locations[j], locations[i]
					}
				}
			}
		}
	}
}

// GetLocationFilterOptions extracts unique values for location filter options
func GetLocationFilterOptions(locations []models.ItemLocation) map[string][]string {
	filterValues := map[string][]string{
		"type":      {},
		"dimension": {},
	}
	
	typeMap := make(map[string]bool)
	dimensionMap := make(map[string]bool)
	
	for _, location := range locations {
		// Add type
		if location.Type != "" && !typeMap[location.Type] {
			typeMap[location.Type] = true
			filterValues["type"] = append(filterValues["type"], location.Type)
		}
		
		// Add dimension
		if location.Dimension != "" && !dimensionMap[location.Dimension] {
			dimensionMap[location.Dimension] = true
			filterValues["dimension"] = append(filterValues["dimension"], location.Dimension)
		}
	}
	
	return filterValues
}