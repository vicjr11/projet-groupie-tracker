package models


type ItemCharacter struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Species  string   `json:"species"`
	Status   string   `json:"status"`
	Gender   string   `json:"gender"`
	Origin   struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Image    string   `json:"image"`
	Episode  []string `json:"episode"`
	URL      string   `json:"url"`
}



type CharacterDetail struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Gender   string   `json:"gender"`
	OriginInfo   Location `json:"originInfo"`
	Location  Location `json:"location"`
	Image    string   `json:"image"`
	Episodes  []string `json:"episode"`
	Created  string   `json:"created"`
	Url      string   `json:"url"`
	Character string   `json:"character"`
	LocationInfo  string  `json:"locationInfo"`
}


type Character struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
	Url     string   `json:"url"`
	Created string   `json:"created"`
}

func (f FilterCharacter) ToFilterParams() any {
	// Implement this method properly
	// This is just a placeholder implementation
	return map[string]string{
		"name":    f.Name,
		"status":  f.Status,
		"species": f.Species,
		"type":    f.Type,
		"gender":  f.Gender,
	}
}

type ListCharacters struct {
	Info struct {
		Count int    `json:"count"`
		Pages int    `json:"pages"`
		Next  string `json:"next"`
	} `json:"info"`
	Results []Character `json:"results"`
}

type FilterCharacter struct {
	Name    string
	Status  string
	Species string
	Type    string
	Gender  string
}
