package models

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


type FilterEpisode struct {
	Name    string
	Episode string
}

func (f FilterEpisode) ToFilterParams() any {
	panic("unimplemented")
}

type ItemEpisode struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Air_date  string `json:"air_date"`
	Episodes   string `json:"episodes"`
	Character string `json:"character"`
	Url       string `json:"url"`
}
