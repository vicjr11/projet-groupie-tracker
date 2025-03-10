package controllers

import (
	"net/http"
	"projet-groupie-tracker/models"
	"projet-groupie-tracker/services"
	temp "projet-groupie-tracker/templates"
	"strconv"
)

func CharactersControllers(w http.ResponseWriter, r *http.Request) {
	// Get page from query parameter, default to 1
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
	}

	// Get characters for specified page
	listCharacters, statusCode, errors := services.CharacterPage(page)
	if errors != nil {
		http.Error(w, errors.Error(), statusCode)
		return
	}

	// Dans la structure paginationData, ajoutons un champ PageNumbers
paginationData := struct {
    Data        models.ListCharacters
    Page        int
    HasPrev     bool
    HasNext     bool
    PrevPage    int
    NextPage    int
    TotalPage   int
    PageNumbers []int // Ajout de cette ligne
}{
    Data:        listCharacters,
    Page:        page,
    HasPrev:     page > 1,
    HasNext:     page < listCharacters.Info.Pages,
    PrevPage:    page - 1,
    NextPage:    page + 1,
    TotalPage:   listCharacters.Info.Pages,
}

// Générer les numéros de page pour la pagination
pageNumbers := []int{}
// Afficher maximum 5 pages autour de la page courante
startPage := page - 2
if startPage < 1 {
    startPage = 1
}
endPage := startPage + 4
if endPage > listCharacters.Info.Pages {
    endPage = listCharacters.Info.Pages
    // Ajuster la page de départ si on est à la fin
    if endPage - 4 > 0 {
        startPage = endPage - 4
    }
}

// Ajouter des ellipses si nécessaire
if startPage > 1 {
    pageNumbers = append(pageNumbers, 1)
    if startPage > 2 {
        pageNumbers = append(pageNumbers, -1) // -1 représente les "..."
    }
}

// Ajouter les pages du milieu
for i := startPage; i <= endPage; i++ {
    pageNumbers = append(pageNumbers, i)
}

// Ajouter des ellipses à la fin si nécessaire
if endPage < listCharacters.Info.Pages {
    if endPage < listCharacters.Info.Pages-1 {
        pageNumbers = append(pageNumbers, -1) // -1 représente les "..."
    }
    pageNumbers = append(pageNumbers, listCharacters.Info.Pages)
}

paginationData.PageNumbers = pageNumbers

	temp.Temp.ExecuteTemplate(w, "character", paginationData)
}
