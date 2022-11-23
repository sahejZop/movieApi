package models

type ResponseModelForSingleMovie struct {
	Code   int        `json:"code,omitempty"`
	Status string     `json:"status,omitempty"`
	Data   MovieModel `json:"data"`
}

type ResponseModelWithStringData struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Data   string `json:"data,omitempty"`
}

type ResponseModelForListOfMovie struct {
	Code   int          `json:"code,omitempty"`
	Status string       `json:"status,omitempty"`
	Data   []MovieModel `json:"data,omitempty"`
}
