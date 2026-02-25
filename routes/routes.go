package routes

import (
	"io"
	"net/http"
)

func NewPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "This is the people handler.\n")
	})
}

func ListTodos() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
