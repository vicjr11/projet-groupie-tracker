package models

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

func (f FilterCharacter) ToFilterParams() any {
	panic("unimplemented")}


type ListCharacters struct {
	Info struct {
		Count int
		Pages int
		Next  string
	}
	Results []Character
}

type FilterCharacter struct {
	Name    string
	Status  string
	Species string
	Type    string
	Gender  string
}

type ItemCharacter struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Species  string   `json:"species"`
	Status   string   `json:"status"`
	Gender   string   `json:"gender"`
	Origin   string `json:"origin"`
	Location string `json:"location"`
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



