package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet-groupie-tracker/models"
	"sync"
	"time"
)

// Structures existantes
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
	Url      string   `json:"url"`
}


// FavoritesManager gère les favoris des utilisateurs
type FavoritesManager struct {
	CharacterFavorites map[int]CharacterDetail
	EpisodeFavorites   map[int] models.Episode
	LocationFavorites  map[int]map[string]interface{}
	mu                 sync.RWMutex // Pour protéger l'accès concurrent
}

// NewFavoritesManager crée une nouvelle instance de gestionnaire de favoris
func NewFavoritesManager() *FavoritesManager {
	return &FavoritesManager{
		CharacterFavorites: make(map[int]CharacterDetail),
		EpisodeFavorites:   make(map[int]models.Episode),
		LocationFavorites:  make(map[int]map[string]interface{}),
	}
}

// Singleton pour le gestionnaire de favoris
var (
	favoritesInstance *FavoritesManager
	once              sync.Once
)

// GetFavoritesManager retourne l'instance unique du gestionnaire de favoris
func GetFavoritesManager() *FavoritesManager {
	once.Do(func() {
		favoritesInstance = NewFavoritesManager()
	})
	return favoritesInstance
}

// AddCharacterToFavorites ajoute un personnage aux favoris
func (fm *FavoritesManager) AddCharacterToFavorites(characterId int) error {
	// Vérifier si le personnage est déjà dans les favoris
	fm.mu.RLock()
	_, exists := fm.CharacterFavorites[characterId]
	fm.mu.RUnlock()

	if exists {
		return fmt.Errorf("character with ID %d is already in favorites", characterId)
	}

	// Récupérer les détails du personnage
	character, statusCode, err := GetCharacterDetail(characterId)
	if err != nil {
		return fmt.Errorf("failed to add character to favorites: %w (status: %d)", err, statusCode)
	}

	// Ajouter aux favoris
	fm.mu.Lock()
	fm.CharacterFavorites[characterId] = character
	fm.mu.Unlock()

	return nil
}

// RemoveCharacterFromFavorites supprime un personnage des favoris
func (fm *FavoritesManager) RemoveCharacterFromFavorites(characterId int) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	if _, exists := fm.CharacterFavorites[characterId]; !exists {
		return fmt.Errorf("character with ID %d is not in favorites", characterId)
	}

	delete(fm.CharacterFavorites, characterId)
	return nil
}

// GetFavoriteCharacters retourne tous les personnages favoris
func (fm *FavoritesManager) GetFavoriteCharacters() []CharacterDetail {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	characters := make([]CharacterDetail, 0, len(fm.CharacterFavorites))
	for _, character := range fm.CharacterFavorites {
		characters = append(characters, character)
	}

	return characters
}

// IsCharacterFavorite vérifie si un personnage est dans les favoris
func (fm *FavoritesManager) IsCharacterFavorite(characterId int) bool {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	_, exists := fm.CharacterFavorites[characterId]
	return exists
}

// AddEpisodeToFavorites ajoute un épisode aux favoris
func (fm *FavoritesManager) AddEpisodeToFavorites(episodeUrl string) error {
	episode, statusCode, err := GetEpisodeDetails(episodeUrl)
	if err != nil {
		return fmt.Errorf("failed to add episode to favorites: %w (status: %d)", err, statusCode)
	}

	fm.mu.Lock()
	defer fm.mu.Unlock()

	if _, exists := fm.EpisodeFavorites[episode.Id]; exists {
		return fmt.Errorf("episode with ID %d is already in favorites", episode.Id)
	}

	fm.EpisodeFavorites[episode.Id] = episode
	return nil
}

// RemoveEpisodeFromFavorites supprime un épisode des favoris
func (fm *FavoritesManager) RemoveEpisodeFromFavorites(episodeId int) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	if _, exists := fm.EpisodeFavorites[episodeId]; !exists {
		return fmt.Errorf("episode with ID %d is not in favorites", episodeId)
	}

	delete(fm.EpisodeFavorites, episodeId)
	return nil
}

// GetFavoriteEpisodes retourne tous les épisodes favoris
func (fm *FavoritesManager) GetFavoriteEpisodes() []models.Episode {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	episodes := make([]models.Episode, 0, len(fm.EpisodeFavorites))
	for _, episode := range fm.EpisodeFavorites {
		episodes = append(episodes, episode)
	}

	return episodes
}

// AddLocationToFavorites ajoute un lieu aux favoris
func (fm *FavoritesManager) AddLocationToFavorites(locationUrl string) error {
	location, statusCode, err := GetLocationDetails(locationUrl)
	if err != nil {
		return fmt.Errorf("failed to add location to favorites: %w (status: %d)", err, statusCode)
	}

	locationId, ok := location["id"].(float64)
	if !ok {
		return fmt.Errorf("failed to extract location ID")
	}

	fm.mu.Lock()
	defer fm.mu.Unlock()

	if _, exists := fm.LocationFavorites[int(locationId)]; exists {
		return fmt.Errorf("location with ID %d is already in favorites", int(locationId))
	}

	fm.LocationFavorites[int(locationId)] = location
	return nil
}

// RemoveLocationFromFavorites supprime un lieu des favoris
func (fm *FavoritesManager) RemoveLocationFromFavorites(locationId int) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	if _, exists := fm.LocationFavorites[locationId]; !exists {
		return fmt.Errorf("location with ID %d is not in favorites", locationId)
	}

	delete(fm.LocationFavorites, locationId)
	return nil
}

// GetFavoriteLocations retourne tous les lieux favoris
func (fm *FavoritesManager) GetFavoriteLocations() []map[string]interface{} {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	locations := make([]map[string]interface{}, 0, len(fm.LocationFavorites))
	for _, location := range fm.LocationFavorites {
		locations = append(locations, location)
	}

	return locations
}

// GetAllFavorites retourne tous les favoris par catégorie
func (fm *FavoritesManager) GetAllFavorites() map[string]interface{} {
	return map[string]interface{}{
		"characters": fm.GetFavoriteCharacters(),
		"episodes":   fm.GetFavoriteEpisodes(),
		"locations":  fm.GetFavoriteLocations(),
	}
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
func GetEpisodeDetails(url string) (models.Episode, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return models.Episode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisodeDetails - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return models.Episode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisodeDetails - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.Episode{}, res.StatusCode, fmt.Errorf("GetEpisodeDetails - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data models.Episode
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return models.Episode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisodeDetails - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

// GetCharacterEpisodes fetches all episodes for a character
func GetCharacterEpisodes(episodeURLs []string) ([]models.Episode, int, error) {
	episodes := make([]models.Episode, 0, len(episodeURLs))

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