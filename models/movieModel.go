package models

import (
	"encoding/json"
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

func ReadMovies() (mapOfMovies map[int]MovieModel, err error) {
	jsonData, err := os.ReadFile("./data/MoviesList.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &mapOfMovies)
	if err != nil {
		return nil, err
	}

	return
}

func GetMoviesFromDatabase() (moviesList []MovieModel, err error) {
	mapOfMovies, err := ReadMovies()
	if err != nil {
		return nil, err
	}
	for index := range mapOfMovies {
		moviesList = append(moviesList, mapOfMovies[index])
	}
	return
}

func ConvertSliceToMap(moviesList []MovieModel) (mapOfMovies map[int]MovieModel) {
	mapOfMovies = map[int]MovieModel{}
	for _, movie := range moviesList {
		mapOfMovies[movie.Id] = movie
	}
	return
}

func SetMovies(mapOfMovies map[int]MovieModel) error {
	jsonData, err := json.Marshal(mapOfMovies)
	if err != nil {
		return err
	}
	err = os.WriteFile("./data/MoviesList.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMovieById(id int) error {
	mapOfMovies, err := ReadMovies()
	if err != nil {
		return err
	}
	delete(mapOfMovies, id)
	SetMovies(mapOfMovies)
	return nil
}

func PutMovieById(id int, movieStruct MovieModel) error {
	mapOfMovies, err := ReadMovies()
	if err != nil {
		return err
	}
	delete(mapOfMovies, id)
	mapOfMovies[movieStruct.Id] = movieStruct
	SetMovies(mapOfMovies)
	return nil
}
