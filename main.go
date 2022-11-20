package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
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
	http.ListenAndServe(":8080", muxRouter)
}

func handleMovies(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		{
			getMovies(writer)
		}
	case "POST":
		{
			postMovies(request)
		}
	}
}

func getMovies(writer http.ResponseWriter) {
	dataToShow, err := json.Marshal(models.GetMovies())
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}

func postMovies(request *http.Request) {
	newData, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var updatedStruct []models.MovieModel
	json.Unmarshal(newData, &updatedStruct)
	models.UpdateMovies(updatedStruct)
}

func handleMovieById(writer http.ResponseWriter, request *http.Request) {
	var id = 0
	for _, val := range mux.Vars(request) {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
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
		movieJson, err := json.Marshal(movie)
		if err != nil {
			panic(err)
		}
		writer.Write(movieJson)
	} else {
		writer.Write([]byte("Movie with this id does not exist"))
		writer.WriteHeader(405)
	}
}

func deleteMovieById(id int, writer http.ResponseWriter) {
	movieMap := models.ReadMovies()
	_, itExists := movieMap[id]
	if itExists {
		models.DeleteMovieById(id)
	} else {
		writer.Write([]byte("Movie with this id does not exist"))
		writer.WriteHeader(405)
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
		writer.Write([]byte("Bad request"))
		writer.WriteHeader(405)
		return
	}
	models.PutMovieById(movieStruct)
}
