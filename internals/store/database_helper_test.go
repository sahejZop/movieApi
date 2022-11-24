package store

import (
	"errors"
	"movieApi/internals/models"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSetMoviesWithoutId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	testCases := []TestCases{
		{
			Input: []models.MovieModelWithoutId{
				{
					Name:     "movie",
					Genre:    "film",
					Rating:   0,
					Plot:     "cinematic video",
					Released: false,
				},
			},
			ExpectedError: nil,
			SqlQuery:      "INSERT INTO movies (Name, Genre, Rating, Plot, Released) VALUES (?, ?, ?, ?, ?)",
		},
		{
			Input: []models.MovieModelWithoutId{
				{
					Name:     "movie",
					Genre:    "film",
					Rating:   0,
					Plot:     "cinematic video",
					Released: false,
				},
			},
			SqlQuery:      "junk query",
			ExpectedError: errors.New("error in sql query"),
		},
	}

	storePtr := New(db)

	for _, v := range testCases {

		if v.ExpectedError == nil {
			for i := 0; i < len(v.Input); i++ {
				mock.ExpectExec(regexp.QuoteMeta(v.SqlQuery)).
					WithArgs(v.Input[i].Name, v.Input[i].Genre, v.Input[i].Rating, v.Input[i].Plot, v.Input[i].Released).
					WillReturnResult(sqlmock.NewResult(int64(i), 1))
			}
		} else {
			for i := 0; i < len(v.Input); i++ {
				mock.ExpectExec(regexp.QuoteMeta(v.SqlQuery)).
					WithArgs(v.Input[i].Name, v.Input[i].Genre, v.Input[i].Rating, v.Input[i].Plot, v.Input[i].Released).
					WillReturnError(v.ExpectedError)
			}
		}

		_, err := storePtr.SetMoviesWithoutId(v.Input)
		if err != nil && !reflect.DeepEqual(err, v.ExpectedError) {
			t.Errorf("error")
		}
	}
}

type TestCases struct {
	Input         []models.MovieModelWithoutId
	SqlQuery      string
	ExpectedError error
}
