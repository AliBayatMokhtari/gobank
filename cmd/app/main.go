package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()

	router.Get("/", handleIndex)

	http.ListenAndServe(":8080", router)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
