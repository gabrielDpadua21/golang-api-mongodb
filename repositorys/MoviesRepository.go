package repositorys

import (
	"log"
	"mongodb-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *MoviesRepository) Connect() {
	session, err := mgo.Dial(m.Server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

func (m *MoviesRepository) GetAll() ([]models.Movie, error) {
	var movies []models.Movie

	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)

	return movies, err
}

func (m *MoviesRepository) GetByID(id string) (models.Movie, error) {
	var movie models.Movie

	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)

	return movie, err
}

func (m *MoviesRepository) Create(movie models.Movie) error {
	err := db.C(COLLECTION).Insert(&movie)

	return err
}

func (m *MoviesRepository) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))

	return err
}

func (m *MoviesRepository) Update(id string, movie models.Movie) error {
	log.Println("id:", id)
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &movie)

	return err
}
