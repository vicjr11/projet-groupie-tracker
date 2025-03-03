package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet-groupie-tracker/models"
	"time"
)



func Getcharacter() (models.ListCharacters, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "https://rickandmortyapi.com/api/character", nil)
	if reqErr != nil {
		return models.ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return models.ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.ListCharacters{}, res.StatusCode, fmt.Errorf("GetCharacter - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data models.ListCharacters
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return models.ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

func CharacterPage(page int) (models.ListCharacters, int, error) {
    // Modify your API URL to include page parameter
    url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", page)
    
    _client := http.Client{
        Timeout: 5 * time.Second,
    }

    req, reqErr := http.NewRequest(http.MethodGet, url, nil)
    if reqErr != nil {
        return models.ListCharacters{}, http.StatusInternalServerError, 
               fmt.Errorf("CharacterPage - Error preparing request: %s", reqErr)
    }

    res, resErr := _client.Do(req)
    if resErr != nil {
        return models.ListCharacters{}, http.StatusInternalServerError, 
               fmt.Errorf("CharacterPage - Error sending request: %s", resErr)
    }

    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return models.ListCharacters{}, res.StatusCode, 
               fmt.Errorf("CharacterPage - Error in response code: %d, message: %s", 
                         res.StatusCode, res.Status)
    }

    var data models.ListCharacters
    decodeErr := json.NewDecoder(res.Body).Decode(&data)
    if decodeErr != nil {
        return models.ListCharacters{}, http.StatusInternalServerError, 
               fmt.Errorf("CharacterPage - Error decoding data: %s", decodeErr.Error())
    }

    return data, res.StatusCode, nil
}
