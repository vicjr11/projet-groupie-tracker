package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet-groupie-tracker/models"
	"time"
	"strings"
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



// EpisodeFilterOptions contains all possible filter criteria for episodes
type EpisodeFilterOptions struct {
	Season    string
	SortBy    string
	SortOrder string
}

// FilterEpisodes filters a list of episodes based on provided criteria
func FilterEpisodes(episodes []models.ItemEpisode, options EpisodeFilterOptions) []models.ItemEpisode {
	if options.Season == "" {
		// No filters applied, but we may still need to sort
		if options.SortBy != "" {
			SortEpisodes(episodes, options.SortBy, options.SortOrder)
		}
		return episodes
	}
	
	filtered := []models.ItemEpisode{}
	
	for _, episode := range episodes {
		// Season filter (assuming episode code format "SxxExx")
		if options.Season != "" {
			seasonCode := "S" + options.Season
			if !strings.Contains(episode.Episodes, seasonCode) {
				continue
			}
		}
		
		// Episode passed all filters
		filtered = append(filtered, episode)
	}
	
	// Apply sorting if specified
	if options.SortBy != "" {
		SortEpisodes(filtered, options.SortBy, options.SortOrder)
	}
	
	return filtered
}

// SortEpisodes sorts a list of episodes based on the specified field and order
func SortEpisodes(episodes []models.ItemEpisode, sortBy, sortOrder string) {
	switch sortBy {
	case "name":
		if sortOrder == "desc" {
			// Sort by name descending
			for i := 0; i < len(episodes)-1; i++ {
				for j := i + 1; j < len(episodes); j++ {
					if episodes[i].Name < episodes[j].Name {
						episodes[i], episodes[j] = episodes[j], episodes[i]
					}
				}
			}
		} else {
			// Sort by name ascending (default)
			for i := 0; i < len(episodes)-1; i++ {
				for j := i + 1; j < len(episodes); j++ {
					if episodes[i].Name > episodes[j].Name {
						episodes[i], episodes[j] = episodes[j], episodes[i]
					}
				}
			}
		}
	case "episode":
		if sortOrder == "desc" {
			// Sort by episode code descending
			for i := 0; i < len(episodes)-1; i++ {
				for j := i + 1; j < len(episodes); j++ {
					if episodes[i].Episodes < episodes[j].Episodes {
						episodes[i], episodes[j] = episodes[j], episodes[i]
					}
				}
			}
		} else {
			// Sort by episode code ascending (default)
			for i := 0; i < len(episodes)-1; i++ {
				for j := i + 1; j < len(episodes); j++ {
					if episodes[i].Episodes > episodes[j].Episodes {
						episodes[i], episodes[j] = episodes[j], episodes[i]
					}
				}
			}
		}
	case "air_date":
		if sortOrder == "desc" {
			// Sort by air date descending
			for i := 0; i < len(episodes)-1; i++ {
				for j := i + 1; j < len(episodes); j++ {
					if episodes[i].Air_date < episodes[j].Air_date {
						episodes[i], episodes[j] = episodes[j], episodes[i]
					}
				}
			}
		} else {
			// Sort by air date ascending (default)
			for i := 0; i < len(episodes)-1; i++ {
				for j := i + 1; j < len(episodes); j++ {
					if episodes[i].Air_date > episodes[j].Air_date {
						episodes[i], episodes[j] = episodes[j], episodes[i]
					}
				}
			}
		}
	}
}

// GetEpisodeFilterOptions extracts unique values for episode filter options
func GetEpisodeFilterOptions(episodes []models.ItemEpisode) map[string][]string {
	filterValues := map[string][]string{
		"season": {},
	}
	
	seasonMap := make(map[string]bool)
	
	for _, episode := range episodes {
		// Extract season number from episode code (e.g., S01E01 -> 01)
		if len(episode.Episodes) >= 4 {
			seasonCode := episode.Episodes[1:3] // Extract "01" from "S01..."
			if !seasonMap[seasonCode] {
				seasonMap[seasonCode] = true
				filterValues["season"] = append(filterValues["season"], seasonCode)
			}
		}
	}
	
	return filterValues
}