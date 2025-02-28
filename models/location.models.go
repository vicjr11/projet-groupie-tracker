package models

type Location struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	Residents string `json:"resident"`
	Url       string  `json:"url"`

}


type ListLocation struct {
	Info struct {
		Count int
		Pages int
		Next  string
	}
	Results []Location
}

type FilterLocation struct {
	Name      string
	Type      string
	Dimension string
}

type ItemLocation struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	Residents string `json:"image"`
	Url       string `json:"url"`
}

func (f FilterLocation) ToFilterParams() any {
	panic("unimplemented")}