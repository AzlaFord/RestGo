package routes

import (
	"fmt"
	"net/http"
)

func NewPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the people handler.")
	})
}
