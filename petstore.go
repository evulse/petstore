package main

import (
	"github.com/evulse/petstore/handlers"
	"log"
	"net/http"
)

func main() {
	petHandler := handlers.PetHandler{}
	indexHandler := handlers.IndexHandler{}
	http.HandleFunc("/pet/", petHandler.DefaultHandler)
	http.HandleFunc("/", indexHandler.DefaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}