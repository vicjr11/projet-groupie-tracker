package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Character struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Origin  struct {
		Name string `json:"name"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
	Url     string  `json:"url"`
}

type ListCharacters struct {
	Info struct {
		Count int
		Pages int
		Next  string
	}
	Results []Character
}

func Getcharacter() (ListCharacters, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "https://rickandmortyapi.com/api/character", nil)
	if reqErr != nil {
		return ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ListCharacters{}, res.StatusCode, fmt.Errorf("GetCharacter - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data ListCharacters
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

type ItemCharacter struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"type"`
	Url     string `json:"url"`
	Image   string `json:"image"`
}

type ResultsCharacter struct {
	Info struct {
		Next string `json:"next"`
	} `json:'info'`
	Results []ItemCharacter `json:'results'`
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



// GetcharacterPage returns pour un characters pour une specifique 
func GetcharacterPage(page int) (ListCharacters, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", page), nil)
	if reqErr != nil {
		return ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ListCharacters{}, res.StatusCode, fmt.Errorf("GetCharacter - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data ListCharacters
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return ListCharacters{}, http.StatusInternalServerError, fmt.Errorf("GetCharacter - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}
