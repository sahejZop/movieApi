package handler

import (
	"encoding/json"
	"io"
	models2 "movieApi/internals/models"
	"net/http"
)

type StoreHandler interface {
	SetMoviesWithoutId(moviesList []models2.MovieModelWithoutId) ([]models2.MovieModel, error)
}

type handler struct {
	st StoreHandler
}

func New(s StoreHandler) *handler {
	return &handler{s}
}

//
//func NotAllowed(writer http.ResponseWriter, _ *http.Request) {
//	WriteErrorResponse(writer, 405, "Method not allowed")
//}
//
//func NotFound(writer http.ResponseWriter, _ *http.Request) {
//	WriteErrorResponse(writer, 404, "Page not found")
//}
//
//func HandleMovies(writer http.ResponseWriter, request *http.Request) {
//	switch request.Method {
//	case "GET":
//		{
//			err := handleGetMovies(writer)
//			if err != nil {
//				WriteInternalServerError(writer)
//			}
//		}
//	case "POST":
//		{
//			err := postMovies(writer, request)
//			if err != nil {
//				WriteInternalServerError(writer)
//			}
//		}
//	}
//}
//
//func handleGetMovies(writer http.ResponseWriter) error {
//	moviesList, err := store.GetMoviesFromDatabase()
//	if err != nil {
//		return err
//	}
//	WriteSuccessResponse(writer, 200, "SUCCESS", moviesList)
//	return nil
//}

func (h *handler) PostMovies(writer http.ResponseWriter, request *http.Request) {
	newData, err := io.ReadAll(request.Body)

	var updatedStruct []models2.MovieModelWithoutId
	err = json.Unmarshal(newData, &updatedStruct)
	if err != nil {
		WriteErrorResponse(writer, 400, "Incorrect format of data")
		return
	}
	moviesStructWithId, err := h.st.SetMoviesWithoutId(updatedStruct)
	if err != nil {
		return
	}
	WriteSuccessResponse(writer, 200, "SUCCESS", moviesStructWithId)
	return
}

//func (*handler) PostMovies(writer http.ResponseWriter, request *http.Request) error {
//	newData, err := io.ReadAll(request.Body)
//	if err != nil {
//		return err
//	}
//	var updatedStruct []models2.MovieModelWithoutId
//	err = json.Unmarshal(newData, &updatedStruct)
//	if err != nil {
//		WriteErrorResponse(writer, 400, "Incorrect format of data")
//		return nil
//	}
//	moviesStructWithId, err := store.SetMoviesWithoutId(updatedStruct)
//	if err != nil {
//		return err
//	}
//	WriteSuccessResponse(writer, 200, "SUCCESS", moviesStructWithId)
//	return nil
//}

//
//func HandleMovieById(writer http.ResponseWriter, request *http.Request) {
//	var id = 0
//	for _, val := range mux.Vars(request) {
//		i, err := strconv.Atoi(val)
//		if err != nil {
//			WriteErrorResponse(writer, 400, "id should be of type int")
//			return
//		}
//		id = i
//	}
//
//	switch request.Method {
//	case "GET":
//		{
//			err := getMovieById(id, writer)
//			if err != nil {
//				WriteInternalServerError(writer)
//			}
//		}
//	case "PUT":
//		{
//			err := putMovieById(id, writer, request.Body)
//			if err != nil {
//				WriteInternalServerError(writer)
//			}
//		}
//	case "DELETE":
//		{
//			err := deleteMovieById(id, writer)
//			if err != nil {
//				WriteInternalServerError(writer)
//			}
//		}
//	}
//
//}
//
//func getMovieById(id int, writer http.ResponseWriter) error {
//	movieMap, err := store.ReadMovies()
//	if err != nil {
//		return err
//	}
//	movie, itExists := movieMap[id]
//	if itExists {
//		WriteSuccessResponseSingleMovie(writer, 200, "SUCCESS", movie)
//	} else {
//		WriteErrorResponse(writer, 404, "Movie with this id does not exist")
//	}
//	return nil
//}
//
//func deleteMovieById(id int, writer http.ResponseWriter) error {
//	movieMap, err := store.ReadMovies()
//	if err != nil {
//		return err
//	}
//	_, itExists := movieMap[id]
//	if itExists {
//		WriteSuccessResponseForDelete(writer, 200, "SUCCESS", "Movie successfully deleted")
//		err := store.DeleteMovieById(id)
//		if err != nil {
//			return err
//		}
//	} else {
//		WriteErrorResponse(writer, 404, "Movie with this id does not exist")
//	}
//	return nil
//}
//
//func putMovieById(id int, writer http.ResponseWriter, body io.ReadCloser) error {
//	movieJson, err := io.ReadAll(body)
//	if err != nil {
//		return err
//	}
//	var movieStruct models2.MovieModel
//	err = json.Unmarshal(movieJson, &movieStruct)
//	if err != nil {
//		WriteErrorResponse(writer, 404, "Bad request")
//		return nil
//	} else {
//		WriteSuccessResponseSingleMovie(writer, 200, "SUCCESS", movieStruct)
//	}
//	err = store.PutMovieById(id, movieStruct)
//	if err != nil {
//		return err
//	}
//	return nil
//}
