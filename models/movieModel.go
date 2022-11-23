package models

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MovieModel struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Genre    string  `json:"genre,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Plot     string  `json:"plot,omitempty"`
	Released bool    `json:"released,omitempty"`
}

var db, err = sql.Open("mysql", "root:jqry4698@tcp(172.17.0.2:3306)/movies_database")

func init() {
	if err != nil {
		fmt.Println(err)
	}
}

func ReadMovies() (mapOfMovies map[int]MovieModel, err error) {
	rows, err := db.QueryContext(context.Background(), "select * from movies")
	m := MovieModel{}
	mapOfMovies = map[int]MovieModel{}
	for rows.Next() {
		rows.Scan(&m.Id, &m.Name, &m.Genre, &m.Rating, &m.Plot, &m.Released)
		mapOfMovies[m.Id] = m
	}

	fmt.Println(mapOfMovies)
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
	for key := range mapOfMovies {
		db.ExecContext(context.Background(), "INSERT INTO movies VALUES (?, ?, ?, ?, ?, ?)",
			mapOfMovies[key].Id, mapOfMovies[key].Name, mapOfMovies[key].Genre, mapOfMovies[key].Rating,
			mapOfMovies[key].Plot, mapOfMovies[key].Released)
	}
	return nil
}

func DeleteMovieById(id int) error {
	_, err = db.ExecContext(context.Background(), "DELETE FROM movies WHERE id = ?", id)
	if err != nil {
		return err
	}
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
