package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time" 

)


    type Episode struct{
	Id int `json:"id"`
	Air_date string `json:"air_date"`
	Episode string `json:"episode"`
	}
	

	type ListEpisode struct{
		Info struct{
			Count int
			Pages int
			Next string
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