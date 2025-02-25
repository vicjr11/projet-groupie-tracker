package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Episode struct {
	Id       int    `json:"id"`
	Air_date string `json:"air_date"`
	Episode  string `json:"episode"`
}

type ListEpisode struct {
	Info struct {
		Count int
		Pages int
		Next  string
	}
	Results []Episode
}

func Getepisode() (ListEpisode, int, error) {
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "https://rickandmortyapi.com/api/episode", nil)
	if reqErr != nil {
		return ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error preparing request: %s", reqErr)
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		return ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error sending request: %s", resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ListEpisode{}, res.StatusCode, fmt.Errorf("GetEpisode - Error in response code: %d, message: %s", res.StatusCode, res.Status)
	}

	var data ListEpisode
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return ListEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error decoding data: %s", decodeErr.Error())
	}

	return data, res.StatusCode, nil
}

type ItemEpisode struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Air_date  string `json:"air_date"`
	Episode   string `json:"episode"`
	Character string `json:"character"`
	Url       string `json:"url"`
}

type ResultsEpisode struct {
	Info struct {
		Next string `json:"next"`
	} `json:'info'`
	Results []ItemEpisode `json:'results'`
}

func GetAllepisode() ([]ItemEpisode, int, error) {
	var ListEpisode []ItemEpisode

	_client := http.Client{
		Timeout: 5 * time.Second,
	}
	page := 1
	for {
		req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rickandmortyapi.com/api/episode?page=%d", page), nil)
		if reqErr != nil {
			return []ItemEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error preparing request: %s", reqErr)
		}

		res, resErr := _client.Do(req)
		if resErr != nil {
			return []ItemEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpidose - Error sending request: %s", resErr)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return []ItemEpisode{}, res.StatusCode, fmt.Errorf("GetEpisode - Error in response code: %d, message: %s", res.StatusCode, res.Status)
		}

		var data ResultsEpisode
		decodeErr := json.NewDecoder(res.Body).Decode(&data)
		if decodeErr != nil {
			return []ItemEpisode{}, http.StatusInternalServerError, fmt.Errorf("GetEpisode - Error decoding data: %s", decodeErr.Error())
		}
		ListEpisode = append(ListEpisode, data.Results...)
		page++
		if data.Info.Next == "" {
			break
		}
		fmt.Println(data)
	}

	return ListEpisode, http.StatusOK, nil
}


