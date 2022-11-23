package handler

import (
	"encoding/json"
	"movieApi/internals/models"
	"net/http"
)

func WriteErrorResponse(writer http.ResponseWriter, code int, error string) {
	response := models.ResponseModelWithStringData{
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

func WriteSuccessResponse(writer http.ResponseWriter, code int, status string, data []models.MovieModel) {
	response := models.ResponseModelForListOfMovie{
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

func WriteSuccessResponseSingleMovie(writer http.ResponseWriter, code int, status string, data models.MovieModel) {
	response := models.ResponseModelForSingleMovie{
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
	response := models.ResponseModelWithStringData{
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

func WriteInternalServerError(writer http.ResponseWriter) {
	WriteErrorResponse(writer, 500, "Internal Server Error")
}
