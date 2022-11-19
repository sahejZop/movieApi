package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type MovieModel struct {
	Id       int
	Name     string
	Genre    string
	Rating   float64
	Plot     string
	Released bool
}

func GetMovie(isUpdated bool) (moviesList []MovieModel) {
	var fileData []byte
	var err = error(nil)
	if isUpdated {
		fileData, err = os.ReadFile("./data/UpdatedMoviesList.json")
	} else {
		fileData, err = os.ReadFile("./data/MoviesList.json")
	}

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileData, &moviesList)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileData))

	return moviesList
}

func UpdateMovie(data []MovieModel) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./data/UpdatedMoviesList.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}
