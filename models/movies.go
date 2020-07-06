package models

import (
	"context"
	"fmt"
	"log"

	"github.com/LockettPundt/GOserver/db"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// routes "github.com/LockettPundt/GOserver/routes"
	utils "github.com/LockettPundt/GOserver/utils"
)

// Movie struct for DB.
type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitompty"`
	ReleaseDate string             `bson:"releaseDate,omitempty"`
	Rating      string             `bson:"rating,omitempty"`
	Status      string             `bson:"status,omitempty"`
	ActorIDs    []string           `bson:"actorIds,omitempty"`
}

// Movies struct for multiple movies.
type Movies struct {
	Movies []Movie `bson:"movie"`
}

// AddMovie pushes the movie from the the form to the DB.
func AddMovie(title, releaseDate, rating, status string) {
	dbName := utils.GoDotEnvVariable("DB_NAME")
	con, err := db.DBconnect()

	if err != nil {
		log.Fatal(err)
	}

	collection := con.Database(dbName).Collection("movies")

	movieToAdd := Movie{
		Title:       title,
		ReleaseDate: releaseDate,
		Rating:      rating,
		Status:      status,
	}
	insertMovie, err := collection.InsertOne(context.TODO(), movieToAdd)

	fmt.Printf("Here is the added Movie ID %v", insertMovie.InsertedID)

}
