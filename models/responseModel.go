package models

import (
	"encoding/json"
	"net/http"
)

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

func WriteErrorResponse(writer http.ResponseWriter, code int, error string) {
	response := ResponseModelWithStringData{
		Code:   code,
		Status: "ERROR",
		Data:   error,
	}
	responseJson, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		panic(err)
	}
	writer.Write(responseJson)
	writer.WriteHeader(code)
	return
}

func WriteSuccessResponse(writer http.ResponseWriter, code int, status string, data []MovieModel) {
	response := ResponseModelForListOfMovie{
		Code:   code,
		Status: status,
		Data:   data,
	}
	dataToShow, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}

func WriteSuccessResponseSingleMovie(writer http.ResponseWriter, code int, status string, data MovieModel) {
	response := ResponseModelForSingleMovie{
		Code:   code,
		Status: status,
		Data:   data,
	}
	dataToShow, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}

func WriteSuccessResponseForDelete(writer http.ResponseWriter, code int, status string, data string) {
	response := ResponseModelWithStringData{
		Code:   code,
		Status: status,
		Data:   data,
	}
	dataToShow, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}
