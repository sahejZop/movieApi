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
	models.WriteErrorResponse(writer, 405, "Method not allowed")
}

func notFound(writer http.ResponseWriter, _ *http.Request) {
	models.WriteErrorResponse(writer, 404, "Page not found")
}

func handleMovies(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		{
			err := handleGetMovies(writer)
			if err != nil {
				WriteInternalServerError(writer)
			}
		}
	case "POST":
		{
			err := postMovies(writer, request)
			if err != nil {
				WriteInternalServerError(writer)
			}
		}
	}
}

func handleGetMovies(writer http.ResponseWriter) error {
	moviesList, err := models.GetMoviesFromDatabase()
	if err != nil {
		return err
	}
	models.WriteSuccessResponse(writer, 200, "SUCCESS", moviesList)
	return nil
}

func postMovies(writer http.ResponseWriter, request *http.Request) error {
	newData, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}
	var updatedStruct []models.MovieModel
	err = json.Unmarshal(newData, &updatedStruct)
	if err != nil {
		models.WriteErrorResponse(writer, 400, "Incorrect format of data")
		return nil
	} else {
		models.WriteSuccessResponse(writer, 200, "SUCCESS", updatedStruct)
	}
	err = models.SetMovies(models.ConvertSliceToMap(updatedStruct))
	if err != nil {
		return err
	}
	return nil
}

func handleMovieById(writer http.ResponseWriter, request *http.Request) {
	var id = 0
	for _, val := range mux.Vars(request) {
		i, err := strconv.Atoi(val)
		if err != nil {
			models.WriteErrorResponse(writer, 400, "id should be of type int")
			return
		}
		id = i
	}

	switch request.Method {
	case "GET":
		{
			err := getMovieById(id, writer)
			if err != nil {
				WriteInternalServerError(writer)
			}
		}
	case "PUT":
		{
			err := putMovieById(id, writer, request.Body)
			if err != nil {
				WriteInternalServerError(writer)
			}
		}
	case "DELETE":
		{
			err := deleteMovieById(id, writer)
			if err != nil {
				WriteInternalServerError(writer)
			}
		}
	}

}

func getMovieById(id int, writer http.ResponseWriter) error {
	movieMap, err := models.ReadMovies()
	if err != nil {
		return err
	}
	movie, itExists := movieMap[id]
	if itExists {
		models.WriteSuccessResponseSingleMovie(writer, 200, "SUCCESS", movie)
	} else {
		models.WriteErrorResponse(writer, 404, "Movie with this id does not exist")
	}
	return nil
}

func deleteMovieById(id int, writer http.ResponseWriter) error {
	movieMap, err := models.ReadMovies()
	if err != nil {
		return err
	}
	_, itExists := movieMap[id]
	if itExists {
		models.WriteSuccessResponseForDelete(writer, 200, "SUCCESS", "Movie successfully deleted")
		err := models.DeleteMovieById(id)
		if err != nil {
			return err
		}
	} else {
		models.WriteErrorResponse(writer, 404, "Movie with this id does not exist")
	}
	return nil
}

func putMovieById(id int, writer http.ResponseWriter, body io.ReadCloser) error {
	movieJson, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	var movieStruct models.MovieModel
	err = json.Unmarshal(movieJson, &movieStruct)
	if err != nil {
		models.WriteErrorResponse(writer, 404, "Bad request")
		return nil
	} else {
		models.WriteSuccessResponseSingleMovie(writer, 200, "SUCCESS", movieStruct)
	}
	err = models.PutMovieById(id, movieStruct)
	if err != nil {
		return err
	}
	return nil
}

func WriteInternalServerError(writer http.ResponseWriter) {
	models.WriteErrorResponse(writer, 500, "Internal Server Error")
}
