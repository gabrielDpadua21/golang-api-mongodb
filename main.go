package main

import (
	"log"
	"mongodb-api/config"
	"mongodb-api/controllers"
	"mongodb-api/repositorys"
	"net/http"

	"github.com/gorilla/mux"
)

var Repository = repositorys.MoviesRepository{}

var conf = config.Config{}

func init() {
	conf.Read()

	Repository.Server = conf.Server
	Repository.Database = conf.Database
	Repository.Connect()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/movie", controllers.Create).Methods("POST")
	r.HandleFunc("/api/v1/movie/{id}", controllers.Update).Methods("PUT")

	var port = ":3000"

	log.Println("Server runing...")
	log.Fatal(http.ListenAndServe(port, r))
}
