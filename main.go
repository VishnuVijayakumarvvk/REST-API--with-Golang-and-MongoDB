package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VISHNUVIJAYAKUMAR/BuildAPI/router"
)

func main() {
	fmt.Println("Welcome to the MongoAPI")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":9500", r))

	fmt.Println("Connection ended.....")
}
