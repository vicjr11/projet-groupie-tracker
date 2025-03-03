package models


type ItemLocation struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	Residents string `json:"image"`
	Url       string `json:"url"`
}



type Location struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Dimension string   `json:"dimension"`
	Residents []string `json:"residents"`
	Url       string   `json:"url"`
}

type ListLocation struct {
	Info struct {
		Count int    `json:"count"`
		Pages int    `json:"pages"`
		Next  string `json:"next"`
	} `json:"info"`
	Results []Location `json:"results"`
}

type FilterLocation struct {
	Name      string
	Type      string
	Dimension string
}

func (f FilterLocation) ToFilterParams() any {
	// Implement this method properly
	// This is just a placeholder implementation
	return map[string]string{
		"name":      f.Name,
		"type":      f.Type,
		"dimension": f.Dimension,
	}
}