package store

import (
	"context"
	"database/sql"
	"movieApi/internals/models"

	_ "github.com/go-sql-driver/mysql"
)

//var db, err = sql.Open("mysql", "root:jqry4698@tcp(172.17.0.2:3306)/movies_database")
//
//func init() {
//	if err != nil {
//		fmt.Println(err)
//	}
//}

type Store struct {
	Db *sql.DB
}

func (store Store) SetMoviesWithoutId(moviesList []models.MovieModelWithoutId) ([]models.MovieModel, error) {
	var moviesStructWithId []models.MovieModel
	for _, val := range moviesList {
		rows, err := store.Db.ExecContext(context.Background(), "INSERT INTO movies (Name, Genre, Rating, Plot, Released) VALUES (?, ?, ?, ?, ?)",
			val.Name, val.Genre, val.Rating, val.Plot, val.Released,
		)
		if err != nil {
			return nil, err
		}
		lastInsertedId, _ := rows.LastInsertId()
		moviesStructWithId = append(moviesStructWithId, models.MovieModel{
			Id:       int(lastInsertedId),
			Name:     val.Name,
			Genre:    val.Genre,
			Rating:   val.Rating,
			Plot:     val.Plot,
			Released: val.Released,
		})
	}

	return moviesStructWithId, nil
}

//	func ReadMovies() (mapOfMovies map[int]models.MovieModel, err error) {
//		rows, err := db.QueryContext(context.Background(), "select * from movies")
//		m := models.MovieModel{}
//		mapOfMovies = map[int]models.MovieModel{}
//		for rows.Next() {
//			rows.Scan(&m.Id, &m.Name, &m.Genre, &m.Rating, &m.Plot, &m.Released)
//			mapOfMovies[m.Id] = m
//		}
//
//		fmt.Println(mapOfMovies)
//		return
//	}
//
//	func GetMoviesFromDatabase() (moviesList []models.MovieModel, err error) {
//		mapOfMovies, err := ReadMovies()
//		if err != nil {
//			return nil, err
//		}
//		for index := range mapOfMovies {
//			moviesList = append(moviesList, mapOfMovies[index])
//		}
//		return
//	}
//
//	func SetMovies(mapOfMovies map[int]models.MovieModel) error {
//		for key := range mapOfMovies {
//			db.ExecContext(context.Background(), "INSERT INTO movies VALUES (?, ?, ?, ?, ?, ?)",
//				mapOfMovies[key].Id, mapOfMovies[key].Name, mapOfMovies[key].Genre, mapOfMovies[key].Rating,
//				mapOfMovies[key].Plot, mapOfMovies[key].Released)
//		}
//		return nil
//	}
//
//func DeleteMovieById(id int) error {
//	_, err = db.ExecContext(context.Background(), "DELETE FROM movies WHERE id = ?", id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func PutMovieById(id int, movieStruct models.MovieModel) error {
//	mapOfMovies, err := ReadMovies()
//	if err != nil {
//		return err
//	}
//	delete(mapOfMovies, id)
//	mapOfMovies[movieStruct.Id] = movieStruct
//	err = SetMovies(mapOfMovies)
//	if err != nil {
//		return err
//	}
//	return nil
//}
