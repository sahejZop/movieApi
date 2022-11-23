package main

import (
	"bytes"
	"encoding/json"
	"io"
	models2 "movieApi/internals/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

type TestCaseListOfMovie struct {
	Input          []models2.MovieModel
	ExpectedOutput models2.ResponseModelForListOfMovie
}

type TestCaseStringOutput struct {
	Input          string
	ExpectedOutput models2.ResponseModelWithStringData
}

func TestGetMovieById(t *testing.T) {
	testCaseSlice := []TestCaseStringOutput{{
		ExpectedOutput: models2.ResponseModelWithStringData{Code: 404, Status: "ERROR", Data: "Movie with this id does not exist"},
		Input:          "0",
	}}

	for i := range testCaseSlice {
		testIdInt, _ := strconv.Atoi(testCaseSlice[i].Input)
		resp := httptest.NewRecorder()
		err := getMovieById(testIdInt, resp)
		if err != nil {
			t.Errorf("error occured")
		}

		if reflect.DeepEqual(resp, testCaseSlice[i]) {
			t.Errorf("mismatch")
		}
	}
}

func TestGetMovies(t *testing.T) {

	moviesList, err := models2.GetMoviesFromDatabase()
	if err != nil {
		t.Errorf("error in getting from database")
	}
	testCase := models2.ResponseModelForListOfMovie{Code: 200, Status: "SUCCESS", Data: moviesList}

	resp := httptest.NewRecorder()

	err = handleGetMovies(resp)
	if err != nil {
		t.Errorf("error occurred in handlerfunc")
	}

	dataFromResponse, err := io.ReadAll(resp.Result().Body)
	var structFromResponse = models2.ResponseModelForListOfMovie{}
	err = json.Unmarshal(dataFromResponse, &structFromResponse)

	if !reflect.DeepEqual(structFromResponse, testCase) {
		t.Errorf("struct mismatch")
	}
}

func TestPostMovies(t *testing.T) {
	testCasesSlice := []TestCaseListOfMovie{
		{
			Input: []models2.MovieModel{{
				Id:       13,
				Name:     "The Ring",
				Genre:    "Horror",
				Rating:   4.5,
				Plot:     "Tv",
				Released: true,
			},
			},
			ExpectedOutput: models2.ResponseModelForListOfMovie{

				Code:   200,
				Status: "SUCCESS",
				Data: []models2.MovieModel{
					{
						Id:       13,
						Name:     "The Ring",
						Genre:    "Horror",
						Rating:   4.5,
						Plot:     "Tv",
						Released: true,
					},
				},
			},
		},
		{
			Input: nil,
			ExpectedOutput: models2.ResponseModelForListOfMovie{
				Code:   200,
				Status: "SUCCESS",
				Data:   nil,
			},
		},
		//{
		//	Input: []models.MovieModel{
		//		{},
		//	},
		//	ExpectedOutput: models.ResponseModelForListOfMovie{
		//		Code:   200,
		//		Status: "SUCCESS",
		//		Data:   nil,
		//	},
		//},
	}

	testCasesSliceWithDifferentInput := []TestCaseStringOutput{
		{
			Input: "test",
			ExpectedOutput: models2.ResponseModelWithStringData{
				Code:   400,
				Status: "ERROR",
				Data:   "Incorrect format of data",
			},
		},
	}

	for i := range testCasesSlice {
		jsonReqBody, err := json.Marshal(testCasesSlice[i].Input)
		if err != nil {
			return
		}

		response := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/movies", bytes.NewReader(jsonReqBody))

		err = postMovies(response, req)
		if err != nil {
			t.Errorf("mismatch")
		}

		dataFromResponse, err := io.ReadAll(response.Result().Body)
		var structFromResponse = models2.ResponseModelForListOfMovie{}
		err = json.Unmarshal(dataFromResponse, &structFromResponse)
		if err != nil {
			t.Errorf("error while unmarshalling")
		}

		if !reflect.DeepEqual(testCasesSlice[i].ExpectedOutput, structFromResponse) {
			t.Errorf("error in setting up data")
		}
	}

	for i := range testCasesSliceWithDifferentInput {
		jsonReqBody, err := json.Marshal(testCasesSliceWithDifferentInput[i].Input)
		if err != nil {
			return
		}

		response := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/movies", bytes.NewReader(jsonReqBody))

		err = postMovies(response, req)
		if err != nil {
			t.Errorf("mismatch")
		}

		dataFromResponse, err := io.ReadAll(response.Result().Body)
		var structFromResponse = models2.ResponseModelWithStringData{}
		err = json.Unmarshal(dataFromResponse, &structFromResponse)
		if err != nil {
			t.Errorf("error while unmarshalling")
		}

		if !reflect.DeepEqual(testCasesSliceWithDifferentInput[i].ExpectedOutput, structFromResponse) {
			t.Errorf("error in setting up data")
		}
	}
}
