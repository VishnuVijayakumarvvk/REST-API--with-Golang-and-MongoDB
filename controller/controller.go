package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/VISHNUVIJAYAKUMAR/BuildAPI/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://mongoAPI:qwertyqwerty@cluster0.7vovqnr.mongodb.net/?retryWrites=true&w=majority"
const dbname = "Netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	uri := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), uri)
	if err != nil {
		log.Fatal(err)
	}
	collection = (*mongo.Collection)(client.Database(dbname).Collection(colName))
	fmt.Println("Database successfully connected:) ")
}

//Mongo Helper

func insertOneMovie(movie model.Netflix) {
	insert, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie inserted successfully with id: ", insert.InsertedID)
}

func updateOneMovie(id string) {
	movieid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": movieid}
	update := bson.M{"$set": bson.M{"watched": true}}
	updateval, _ := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println("Movie is successfully updated: ", updateval.ModifiedCount)
}

func deleteOneMovie(id string) {
	movieid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": movieid}
	deleteone, _ := collection.DeleteOne(context.Background(), filter)
	fmt.Println("Movie is successfully deleted: ", deleteone.DeletedCount)
}

func deleteAllMovies() int {
	deletemovie, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movie deleted : ", deletemovie.DeletedCount)
	return int(deletemovie.DeletedCount)
}

func getallmovies() []primitive.M {
	cursor, _ := collection.Find(context.Background(), bson.M{})
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

// controller functions

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to cred opertaion of api<h1>"))
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	var movies = getallmovies()
	json.NewEncoder(w).Encode(movies)
}

func Createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")
	var moviess model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&moviess)
	insertOneMovie(moviess)
	json.NewEncoder(w).Encode(moviess)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")
	var deletecout = deleteAllMovies()
	fmt.Println("Movie deleted")
	json.NewEncoder(w).Encode(deletecout)
}

func DeleteoneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Course deleted")
}
