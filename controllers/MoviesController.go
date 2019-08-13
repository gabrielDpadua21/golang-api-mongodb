package controllers

import (
	"encoding/json"
	"log"
	"mongodb-api/models"
	"mongodb-api/repositorys"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var Repo = repositorys.MoviesRepository{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := Repo.GetAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJson(w, http.StatusOK, movies)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movie, err := Repo.GetByID(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}

	respondWithJson(w, http.StatusOK, movie)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request post")
		return
	}

	movie.ID = bson.NewObjectId()

	if err := Repo.Create(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, movie)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request put")
		return
	}

	log.Println(params["id"])

	if err := Repo.Update(params["id"], movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, movie)
}
