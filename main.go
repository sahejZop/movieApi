package main

import (
	"database/sql"
	_ "database/sql"
	"log"
	"movieApi/internals/handler"
	"movieApi/internals/store"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	var db, _ = sql.Open("mysql", "root:jqry4698@tcp(172.17.0.2:3306)/movies_database")
	store := store.Store{
		Db: db,
	}
	httpHandler := handler.New(store)
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/movies", httpHandler.PostMovies)
	//muxRouter.HandleFunc("/movies", handler.HandleMovies).Methods("GET", "POST")
	//muxRouter.HandleFunc("/movies/{id}", handler.HandleMovieById).Methods("GET", "PUT", "DELETE")
	//muxRouter.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	//muxRouter.MethodNotAllowedHandler = http.HandlerFunc(handler.NotAllowed)
	err := http.ListenAndServe(":8080", muxRouter)
	if err != nil {
		log.Fatal(err)
	}
}

//func SetupServer() {
//}
