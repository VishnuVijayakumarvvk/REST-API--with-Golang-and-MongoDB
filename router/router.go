package router

import (
	"github.com/VISHNUVIJAYAKUMAR/BuildAPI/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Home).Methods("GET")
	router.HandleFunc("/api/movies",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.Createmovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteoneCourse).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllCourses).Methods("DELETE")

	return router
}