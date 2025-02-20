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

func Getepisode() (Episode, int, error) {
	// Création d'un client HTTP avec un timeout de 5 secondes stocké dans la variable : _client
	_client := http.Client{
		Timeout: 5 * time.Second,
	}

	// Préparation de la requête grâce à http.NewRequest
	// req : contient la requête
	// reqErr : capture l'erreur lors de la préparation de celle-ci sinon nil
	req, reqErr := http.NewRequest(http.MethodGet, "https://rickandmortyapi.com/api/episode", nil)
	if reqErr != nil {
		// Gestion en cas d'erreur lors de la préparation de celle-ci
		// Retourne une structure vide, le code HTTP 500 puis un message pour expliquer l'erreur
		return Episode{}, http.StatusInternalServerError, fmt.Errorf("Getepisode - Erreur lors de la préparation de la requête : %s", reqErr)
	}

	// Envois de la requête grâce au client : _client
	// req est la requête préparée lors de l'étape précédente
	res, resErr := _client.Do(req)
	if resErr != nil {
		// Gestion en cas d'erreur lors de l'envois de la requête
		// Retourne une structure vide, le code HTTP 500 puis un message pour expliquer l'erreur
		return Episode{}, http.StatusInternalServerError, fmt.Errorf("Getepisode- Erreur lors de l'envois de la requête : %s", reqErr)
	}

	// Fermeture du corps de la réponse différée, pour récupérer l'ensemble des données
	defer res.Body.Close()

	// Vérification de la présence d'une erreur dans la réponse soit code de réponse différent de 200
	if res.StatusCode != http.StatusOK {
		// Gestion en cas d'erreur dans la réponse de la requête
		// Retourne une structure vide, le code de la réponse puis un message pour expliquer l'erreur
		return Episode{}, res.StatusCode, fmt.Errorf("Getepisode - Erreur dans la réponse code : %d, message : %s", res.StatusCode, res.Status)
	}

	// Déclaration de la variable pour stocker les données décodées
	var data Episode

	// Décodage des données de la réponse au format JSON, pour les stocker dans la variable : data
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		// Gestion en cas d'erreur dans la réponse de la requête
		// Retourne une structure vide, le code de la réponse puis un message pour expliquer l'erreur
		return Episode{}, http.StatusInternalServerError, fmt.Errorf("Getepisode- Erreur dans lors du décodage des données : %s", decodeErr.Error())
	}
	// Retourne les données, le code HTTP et nil pour l'erreur
	return data, res.StatusCode, nil
}