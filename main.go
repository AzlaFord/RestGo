package main

import (
	"RestGo/routes"
	"RestGo/storage"
	"log"
	"net/http"
)

func main() {
	storage.ReadStorage()
	mux := http.NewServeMux()
	mux.Handle("/resources", http.NotFoundHandler())
	mux.Handle("/resources/people", routes.NewPeopleHandler())
	mux.Handle("/list", routes.ListTodos)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
