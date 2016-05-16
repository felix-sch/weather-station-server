package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is alive.")

	vars := mux.Vars(r)
	key := vars["key"]

	if len(key) > 0 {
		fmt.Fprintf(w, "\nKey: %s", key)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./app/index.html")
}

func main() {
	r := mux.NewRouter()

	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", ApiHandler).Methods("GET")
	s.HandleFunc("/{key}/", ApiHandler)

	r.HandleFunc("/", HomeHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/")))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
