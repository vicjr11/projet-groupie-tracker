package services

import (
	"projet-groupie-tracker/models"
	"projet-groupie-tracker/utils"
	"strings"
)


// SearchCharacters recherche des personnages par nom et espèce
func SearchCharacters(query string, characters []models.ItemCharacter) []models.SearchResult {

	results := []models.SearchResult{}

	// Convertir la requête en minuscules pour une recherche insensible à la casse
	loweredQuery := strings.ToLower(query)

	for _, character := range characters {
		matching := []string{}

		// Vérifier si le nom correspond
		if strings.Contains(strings.ToLower(character.Name), loweredQuery) {
			matching = append(matching, "name")
		}

		// Vérifier si l'espèce correspond
		if strings.Contains(strings.ToLower(character.Species), loweredQuery) {
			matching = append(matching, "species")
		}

		// Si au moins une propriété correspond, ajouter aux résultats
		if len(matching) > 0 {
			results = append(results, models.SearchResult{
				Type:     "character",
				Data:     character,
				Matching: matching,
			})
		}
	}

	return results
}

// SearchEpisodes recherche des épisodes par nom et code d'épisode
func SearchEpisodes(query string, episodes []models.ItemEpisode) []models.SearchResult {
	results := []models.SearchResult{}

	// Convertir la requête en minuscules pour une recherche insensible à la casse
	loweredQuery := strings.ToLower(query)

	for _, episode := range episodes {
		matching := []string{}

		// Vérifier si le nom correspond
		if strings.Contains(strings.ToLower(episode.Name), loweredQuery) {
			matching = append(matching, "name")
		}

		// Vérifier si le code d'épisode correspond
		if strings.Contains(strings.ToLower(episode.Episodes), loweredQuery) {
			matching = append(matching, "episode")
		}

		// Si au moins une propriété correspond, ajouter aux résultats
		if len(matching) > 0 {
			results = append(results, models.SearchResult{
				Type:     "episode",
				Data:     episode,
				Matching: matching,
			})
		}
	}

	return results
   }

// SearchLocations recherche des lieux par nom et type
func SearchLocations(query string, locations []models.ItemLocation) []models.SearchResult {
	results := []models.SearchResult{}

	// Convertir la requête en minuscules pour une recherche insensible à la casse
	loweredQuery := strings.ToLower(query)

	for _, location := range locations {
		matching := []string{}

		// Vérifier si le nom correspond
		if strings.Contains(strings.ToLower(location.Name), loweredQuery) {
			matching = append(matching, "name")
		}

		// Vérifier si le type correspond
		if strings.Contains(strings.ToLower(location.Type), loweredQuery) {
			matching = append(matching, "type")
		}

		// Vérifier si la dimension correspond (propriété supplémentaire)
		if strings.Contains(strings.ToLower(location.Dimension), loweredQuery) {
			matching = append(matching, "dimension")
		}

		// Si au moins une propriété correspond, ajouter aux résultats
		if len(matching) > 0 {
			results = append(results, models.SearchResult{
				Type:     "location",
				Data:     location,
				Matching: matching,
			})
		}
	}

	return results
}

// SearchAll effectue une recherche sur toutes les collections
func SearchAll(query string) ([]models.SearchResult, error) {
	allResults := []models.SearchResult{}

	// Récupérer tous les personnages
	characters, status, err := utils.GetAllcharacter()
	if err != nil || status != 200 {
		return nil, err
	}

	// Récupérer tous les épisodes
	episodes, status, err := utils.GetAllepisode()
	if err != nil || status != 200 {
		return nil, err
	}

	// Récupérer tous les lieux
	locations, status, err := utils.GetAlllocation()
	if err != nil || status != 200 {
		return nil, err
	}

	// Convertir les utils.ItemCharacter en services.ItemCharacter
	var serviceCharacters []models.ItemCharacter
	for _, char := range characters {
		serviceCharacters = append(serviceCharacters, models.ItemCharacter{
			ID:       char.ID,
			Name:     char.Name,
			Species:  char.Species,
			Status:   char.Status,
			Gender:   char.Gender,
			Origin:   char.Origin,
			Location: char.Location,
			Image:    char.Image,
			Episode:  char.Episode,
			URL:      char.URL,
		})
	}

	// Rechercher dans chaque collection
	characterResults := SearchCharacters(query, serviceCharacters)
	episodeResults := SearchEpisodes(query, episodes)
	locationResults := SearchLocations(query, locations)

	// Combiner tous les résultats
	allResults = append(allResults, characterResults...)
	allResults = append(allResults, episodeResults...)
	allResults = append(allResults, locationResults...)

	return allResults, nil
}
