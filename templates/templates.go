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
	temp, tempErr := template.ParseGlob("./templates/*.html")
	if tempErr != nil {
		fmt.Printf("Erreur Template - Une erreur lors du chargement des template \n message d'erreur : %v\n", tempErr.Error())
		os.Exit(1)
	}
	Temp = temp
}
