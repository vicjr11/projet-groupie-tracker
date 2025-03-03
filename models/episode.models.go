package models



type ItemEpisode struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Air_date  string `json:"air_date"`
	Episodes   string `json:"episodes"`
	Character string `json:"character"`
	Url       string `json:"url"`
}


type Episode struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Air_date  string   `json:"air_date"`
	Episode   string   `json:"episode"`
	Characters []string `json:"characters"`
	Url       string   `json:"url"`
}

type ListEpisode struct {
	Info struct {
		Count int    `json:"count"`
		Pages int    `json:"pages"`
		Next  string `json:"next"`
	} `json:"info"`
	Results []Episode `json:"results"`
}

type FilterEpisode struct {
	Name    string
	Episode string
}

func (f FilterEpisode) ToFilterParams() any {
	// Implement this method properly
	// This is just a placeholder implementation
	return map[string]string{
		"name":    f.Name,
		"episode": f.Episode,
	}
}