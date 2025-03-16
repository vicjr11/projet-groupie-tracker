# Projet Rick et Morty

Une application web permettant d'explorer l'univers de Rick et Morty à travers ses personnages, épisodes et lieux.

## Structure du projet

```
├── assets/
│   ├── css/
│   │   ├── animation.css
│   │   ├── characters.css
│   │   ├── details.css
│   │   ├── episodes.css
│   │   ├── favoris.css
│   │   ├── home.css
│   │   ├── location.css
│   │   ├── search.css
│   │   └── universal.css
│   └── js/
│       ├── character.js
│       ├── details.js
│       └── favoris.js
├── controllers/
│   ├── character.controllers.go
│   ├── episode.controllers.go
│   ├── favoris.controllers.go
│   ├── home.controllers.go
│   ├── location.controllers.go
│   └── search.controllers.go
├── models/
│   ├── character.models.go
│   ├── episode.models.go
│   ├── location.models.go
│   ├── search.models.go
│   └── favoris.models.go
├── routes/
│   ├── main.routes.go
│   └── route.go
├── service/
│   ├── character.services.go
│   ├── episode.services.go
│   ├── favoris.services.go
│   ├── location.service.go
│   └── search.service.go
├── templates/
│   ├── character/
│   ├── details/
│   ├── episode/
│   ├── favoris/
│   ├── home/
│   ├── location/
│   ├── search/
│   └── templates.go
├── utils/
│   └── pagination.go
├── main.go
└── test.json
```

## Description

Cette application est construite selon une architecture MVC (Modèle-Vue-Contrôleur) en Go, permettant d'explorer et de gérer les données de l'univers Rick et Morty.

## Fonctionnalités

- Affichage des personnages, épisodes et lieux
- Recherche avancée
- Système de favoris 
- Navigation paginée
- Vue détaillée des éléments

## Installation

```bash
# Cloner le dépôt
git clone [url-du-repo]

# Accéder au répertoire
cd [nom-du-repo]

# Installer les dépendances
go mod tidy

# Lancer l'application
go run main.go
```

## Utilisation

Accédez à l'application via `http://localhost:8080` dans votre navigateur.

## Développement

Le projet est organisé comme suit:
- `models`: définition des structures de données
- `controllers`: gestion des requêtes et réponses
- `service`: logique métier
- `templates`: fichiers de présentation
- `routes`: définition des points d'entrée API
- `assets`: ressources frontend (CSS/JS)
- `utils`: fonctions utilitaires

## Tests

Exécutez les tests avec `go test ./...`

## lancer le serveur 

on lance le serveur dans le terminal en écrivant go run main.go

