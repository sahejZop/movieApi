package main

import (
	"fmt"
	"movieApi/models"
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
	fmt.Println(models.GetMovie(true))
}
