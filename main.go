package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"movieApi/models"
	"net/http"
)

func main() {
	fmt.Println(models.GetMovie(false))
	models.UpdateMovie([]models.MovieModel{
		{
			1,
			"Breaking bad",
			"Drama",
			9.9,
			"Chemistry professor breaks bad",
			true,
		},
	})
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", handleStuff)
	muxRouter.HandleFunc("/update", handleUpdate)
	http.ListenAndServe(":8080", muxRouter)
}

func handleUpdate(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		newData, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		var updatedStruct []models.MovieModel
		json.Unmarshal(newData, &updatedStruct)
		models.UpdateMovie(updatedStruct)
	} else {
		writer.WriteHeader(405)
		writer.Write([]byte("Method not allowed"))
	}
}

func handleStuff(writer http.ResponseWriter, request *http.Request) {
	dataToShow, err := json.Marshal(models.GetMovie(false))
	if err != nil {
		panic(err)
	}
	writer.Write(dataToShow)
}
