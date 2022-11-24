package store

import models2 "movieApi/internals/models"

type StoreHandler interface {
	SetMoviesWithoutId(moviesList []models2.MovieModelWithoutId) ([]models2.MovieModel, error)
}
