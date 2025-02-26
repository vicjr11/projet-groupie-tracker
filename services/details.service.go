

package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CharacterDetail struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Gender   string   `json:"gender"`
	Origin   Location `json:"origin"`
	Location Location `json:"location"`
	Image    string   `json:"image"`
	Episode  []string `json:"episode"`
	Created  string   `json:"created"`
	Url      string    `json:"url"`
}



// GetCharacterDetail fetches complete information for a single character by ID
func GetCharacterDetail(id int) (CharacterDetail, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", id), nil)
	if reqErr != nil {
		return CharacterDetail{}, http.StatusInternalServerError, fmt.Errorf("GetCharacterDetail - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return CharacterDetail{}, http.StatusInternalServerError, fmt.Errorf("GetCharacterDetail - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CharacterDetail{}, res.StatusCode, fmt.Errorf("GetCharacterDetail - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data CharacterDetail
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return CharacterDetail{}, http.StatusInternalServerError, fmt.Errorf("GetCharacterDetail - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

// GetEpisodeDetails fetches information for a specific episode by URL
func GetEpisodeDetails(url string) (Episode, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return Episode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisodeDetails - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return Episode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisodeDetails - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Episode{}, res.StatusCode, fmt.Errorf("GetEpisodeDetails - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data Episode
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return Episode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisodeDetails - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

// GetCharacterEpisodes fetches all episodes for a character
func GetCharacterEpisodes(episodeURLs []string) ([]Episode, int, error) {
	episodes := make([]Episode, 0, len(episodeURLs))

	for _, url := range episodeURLs {
		episode, statusCode, err := GetEpisodeDetails(url)
		if err != nil {
			return nil, statusCode, fmt.Errorf("GetCharacterEpisodes - Error fetching episode: %s", err)
		}
		episodes = append(episodes, episode)
	}

	return episodes, http.StatusOK, nil
}

// GetLocationDetails fetches details for a location by URL
func GetLocationDetails(url string) (map[string]interface{}, int, error) {
	// If URL is empty, return empty location
	if url == "" {
		return map[string]interface{}{
			"name": "Unknown",
			"type": "Unknown",
			"dimension": "Unknown",
		}, http.StatusOK, nil
	}

	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetLocationDetails - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetLocationDetails - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("GetLocationDetails - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data map[string]interface{}
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("GetLocationDetails - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}