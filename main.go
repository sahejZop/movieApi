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

func notAllowed(writer http.ResponseWriter, _ *http.Request) {
	writeErrorResponse(writer, 405, "Method not allowed")
}

func notFound(writer http.ResponseWriter, request *http.Request) {
	writeErrorResponse(writer, 404, "Page not found")
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
	writeSuccessResponse(writer, 200, "SUCCESS", models.GetMovies())
}

func postMovies(writer http.ResponseWriter, request *http.Request) {
	newData, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var updatedStruct []models.MovieModel
	err = json.Unmarshal(newData, &updatedStruct)
	if err != nil {
		writeErrorResponse(writer, 400, "Incorrect format of data")
		return
	} else {
		writeSuccessResponse(writer, 200, "SUCCESS", updatedStruct)
	}
	models.UpdateMovies(updatedStruct)
}

func handleMovieById(writer http.ResponseWriter, request *http.Request) {
	var id = 0
	for _, val := range mux.Vars(request) {
		i, err := strconv.Atoi(val)
		if err != nil {
			writeErrorResponse(writer, 400, "id should be of type int")
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
		writeSuccessResponseSingleMovie(writer, 200, "SUCCESS", movie)
	} else {
		writeErrorResponse(writer, 404, "Movie with this id does not exist")
	}
}

func deleteMovieById(id int, writer http.ResponseWriter) {
	movieMap := models.ReadMovies()
	_, itExists := movieMap[id]
	if itExists {
		writeSuccessResponseForDelete(writer, 200, "SUCCESS", "Movie successfully deleted")
		models.DeleteMovieById(id)
	} else {
		writeErrorResponse(writer, 404, "Movie with this id does not exist")
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
		writeErrorResponse(writer, 404, "Bad request")
		return
	} else {
		writeSuccessResponseSingleMovie(writer, 200, "SUCCESS", movieStruct)
	}
	models.PutMovieById(movieStruct)
}

func writeErrorResponse(writer http.ResponseWriter, code int, error string) {
	response := models.ResponseModelWithStringData{
		Code:   code,
		Status: "ERROR",
		Data:   error,
	}
	responseJson, err := json.MarshalIndent(response, "", "")
	if err != nil {
		panic(err)
	}
	writer.Write(responseJson)
	writer.WriteHeader(code)
	return
}

func writeSuccessResponse(writer http.ResponseWriter, code int, status string, data []models.MovieModel) {
	response := models.ResponseModelForListOfMovie{
		Code:   code,
		Status: status,
		Data:   data,
	}
	dataToShow, err := json.MarshalIndent(response, "", "")
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}

func writeSuccessResponseSingleMovie(writer http.ResponseWriter, code int, status string, data models.MovieModel) {
	response := models.ResponseModelForSingleMovie{
		Code:   code,
		Status: status,
		Data:   data,
	}
	dataToShow, err := json.MarshalIndent(response, "", "")
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}

func writeSuccessResponseForDelete(writer http.ResponseWriter, code int, status string, data string) {
	response := models.ResponseModelWithStringData{
		Code:   code,
		Status: status,
		Data:   data,
	}
	dataToShow, err := json.MarshalIndent(response, "", "")
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}
