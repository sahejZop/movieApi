package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type MovieModel struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Genre    string  `json:"genre,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Plot     string  `json:"plot,omitempty"`
	Released bool    `json:"released,omitempty"`
}

func ReadMovies() (mapOfMovies map[int]MovieModel) {
	mapOfMovies = map[int]MovieModel{}
	jsonData, err := os.ReadFile("./data/MoviesList.json")
	if err != nil {
		panic(err)
	}
	var movies []MovieModel
	err = json.Unmarshal(jsonData, &movies)
	if err != nil {
		panic(err)
	}

	for _, movie := range movies {
		mapOfMovies[(movie.Id)] = movie
	}

	return
}

func GetMovies() (moviesList []MovieModel) {
	var fileData []byte
	var err = error(nil)
	fileData, err = os.ReadFile("./data/MoviesList.json")

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

func UpdateMovies(data []MovieModel) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./data/MoviesList.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func GetMovieById(id int) (movie MovieModel) {
	mapOfMovies := ReadMovies()

	return mapOfMovies[id]
}

func DeleteMovieById(id int) {
	mapOfMovies := ReadMovies()
	delete(mapOfMovies, id)
	var movies []MovieModel
	for _, movie := range mapOfMovies {
		movies = append(movies, movie)
	}

	UpdateMovies(movies)
}

func PutMovieById(movieStruct MovieModel) {
	mapOfMovies := ReadMovies()

	var movies []MovieModel
	movies = append(movies, movieStruct)
	for _, movie := range mapOfMovies {
		movies = append(movies, movie)
	}
	UpdateMovies(movies)
}
