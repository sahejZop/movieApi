package models

type MovieModel struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Genre    string  `json:"genre,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Plot     string  `json:"plot,omitempty"`
	Released bool    `json:"released,omitempty"`
}

type MovieModelWithoutId struct {
	Name     string  `json:"name,omitempty"`
	Genre    string  `json:"genre,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Plot     string  `json:"plot,omitempty"`
	Released bool    `json:"released,omitempty"`
}
