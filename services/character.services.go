package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time" 

)

type Character struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Species string `json:"species"`
	Type string `json:"type"`
	Origin struct{
		Name string `json:"name"`
	} `json:"origin"`
	Location struct{
		Name string `json:"name"`
	} `json:"location"`
	Image string `json:"image"`
	Episode []string `json:"episode"`
}


type ListCharacters struct{
	Info struct{
		Count int
		Pages int
		Next string
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

