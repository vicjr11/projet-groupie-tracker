package temp

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Variable globale qui sert à stocker les templates chargés
var Temp *template.Template

// Méthode permettant de charger l'ensemble des templates
func InitTemplates() {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	
	// Définir d'abord le FuncMap
	funcMap := template.FuncMap{
		"add":      func(a, b int) int { return a + b },
		"Subtract": func(a, b int) int { return a - b },
		"sequence": generateSequence,
	}
	
	// Créer un template avec le FuncMap
	temp := template.New("").Funcs(funcMap)
	
	// Ensuite parser les templates
	var tempErr error
	temp, tempErr = temp.ParseGlob("./templates/*.html")
	
	if tempErr != nil {
		fmt.Printf("Erreur Template - Une erreur lors du chargement des template \n message d'erreur : %v\n", tempErr.Error())
		os.Exit(1)
	}

	// Assigner à la variable globale
	Temp = temp
}

func generateSequence(start, end int) []int {
	// Limit the number of pages displayed to prevent too many buttons
	maxPages := 5

	if end-start+1 > maxPages {
		if start == 1 {
			// Show first few pages
			end = start + maxPages - 1
		} else {
			// Center around current page
			middle := (start + end) / 2
			start = middle - maxPages/2
			if start < 1 {
				start = 1
			}
			end = start + maxPages - 1
		}
	}

	result := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}