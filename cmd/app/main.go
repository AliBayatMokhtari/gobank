package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()

	router.Get("/", handleIndex)
	router.Get("/user/{id}", HandleGetUser)

	http.ListenAndServe(":8080", router)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:       1,
		Name:     "John Doe",
		Username: "ali_bm",
	}

	data, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Failed to marshal user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(data)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := User{
		ID:       idInt,
		Name:     "John Doe",
		Username: "ali_bm",
	}

	data, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Failed to marshal user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(data)
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"user_name"`
}
