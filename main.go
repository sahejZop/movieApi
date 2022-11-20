package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"movieApi/models"
	"net/http"
	"strconv"
)

func main() {
	setupServer()
}

func setupServer() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/movies", handleMovies).Methods("GET", "POST")
	muxRouter.HandleFunc("/movie/{id}", handleMovieById).Methods("GET", "PUT", "DELETE")
	muxRouter.NotFoundHandler = http.HandlerFunc(notFound)
	muxRouter.MethodNotAllowedHandler = http.HandlerFunc(notAllowed)
	err := http.ListenAndServe(":8080", muxRouter)
	if err != nil {
		log.Fatal(err)
	}
}

func notAllowed(writer http.ResponseWriter, request *http.Request) {
	response := models.ResponseModelForDeleteOrError{
		Code:   405,
		Status: "ERROR",
		Data:   "Method not allowed",
	}
	responseJson, err := json.MarshalIndent(response, "", "")
	if err != nil {
		panic(err)
	}
	writer.Write(responseJson)
	writer.WriteHeader(405)
	return
}

func notFound(writer http.ResponseWriter, request *http.Request) {
	response := models.ResponseModelForDeleteOrError{
		Code:   404,
		Status: "ERROR",
		Data:   "Page not found",
	}
	responseJson, err := json.MarshalIndent(response, "", "")
	if err != nil {
		panic(err)
	}
	writer.Write(responseJson)
	writer.WriteHeader(404)
	return
}

func handleMovies(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		{
			getMovies(writer)
		}
	case "POST":
		{
			postMovies(writer, request)
		}
	}
}

func getMovies(writer http.ResponseWriter) {
	response := models.ResponseModelForListOfMovie{
		Code:   200,
		Status: "SUCCESS",
		Data:   models.GetMovies(),
	}
	dataToShow, err := json.MarshalIndent(response, "", "")
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}

func postMovies(writer http.ResponseWriter, request *http.Request) {
	newData, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var updatedStruct []models.MovieModel
	err = json.Unmarshal(newData, &updatedStruct)
	if err != nil {
		response := models.ResponseModelForDeleteOrError{
			Code:   400,
			Status: "ERROR",
			Data:   "Incorrect format of data",
		}
		responseJson, err := json.MarshalIndent(response, "", "")
		if err != nil {
			panic(err)
		}
		writer.Write(responseJson)
		writer.WriteHeader(400)
		return
	} else {
		response := models.ResponseModelForListOfMovie{
			Code:   200,
			Status: "SUCCESS",
			Data:   updatedStruct,
		}
		responseJson, err := json.MarshalIndent(response, "", "")
		if err != nil {
			panic(err)
		}
		writer.Write(responseJson)
	}
	models.UpdateMovies(updatedStruct)
}

func handleMovieById(writer http.ResponseWriter, request *http.Request) {
	var id = 0
	for _, val := range mux.Vars(request) {
		i, err := strconv.Atoi(val)
		if err != nil {
			response := models.ResponseModelForDeleteOrError{
				Code:   400,
				Status: "ERROR",
				Data:   "Id should be of type int",
			}
			responseJson, err := json.MarshalIndent(response, "", "")
			if err != nil {
				panic(err)
			}
			writer.Write(responseJson)
			writer.WriteHeader(400)
			return
		}
		id = i
	}

	switch request.Method {
	case "GET":
		{
			getMovieById(id, writer)
		}
	case "PUT":
		{
			putMovieById(writer, request.Body)
		}
	case "DELETE":
		{
			deleteMovieById(id, writer)
		}
	}

}

func getMovieById(id int, writer http.ResponseWriter) {
	movieMap := models.ReadMovies()
	movie, itExists := movieMap[id]
	if itExists {
		response := models.ResponseModelForSingleMovie{
			Code:   200,
			Status: "SUCCESS",
			Data:   movie,
		}
		responseJson, err := json.MarshalIndent(response, "", "")
		if err != nil {
			panic(err)
		}
		writer.Write(responseJson)
	} else {
		response := models.ResponseModelForDeleteOrError{
			Code:   404,
			Status: "ERROR",
			Data:   "Movie with this id does not exist",
		}
		responseJson, err := json.MarshalIndent(response, "", "")
		if err != nil {
			panic(err)
		}
		writer.Write(responseJson)
		writer.WriteHeader(404)
	}
}

func deleteMovieById(id int, writer http.ResponseWriter) {
	movieMap := models.ReadMovies()
	_, itExists := movieMap[id]
	if itExists {
		models.DeleteMovieById(id)
	} else {
		response := models.ResponseModelForDeleteOrError{
			Code:   404,
			Status: "ERROR",
			Data:   "Movie with this id does not exist",
		}
		responseJson, err := json.MarshalIndent(response, "", "")
		if err != nil {
			panic(err)
		}
		writer.Write(responseJson)
		writer.WriteHeader(404)
	}
}

func putMovieById(writer http.ResponseWriter, body io.ReadCloser) {
	movieJson, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}
	var movieStruct models.MovieModel
	err = json.Unmarshal(movieJson, &movieStruct)
	if err != nil {
		response := models.ResponseModelForDeleteOrError{
			Code:   404,
			Status: "ERROR",
			Data:   "Bad request",
		}
		responseJson, err := json.MarshalIndent(response, "", "")
		if err != nil {
			panic(err)
		}
		writer.Write(responseJson)
		writer.WriteHeader(400)
		return
	}
	models.PutMovieById(movieStruct)
}
