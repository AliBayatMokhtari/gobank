package main

import (
	"encoding/json"
	"net/http"
	"time"

	dtos "github.com/alibayatmokhtari/gobank/cmd/types/DTOs"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()

	router.Get("/register", registerUser)

	http.ListenAndServe(":8080", router)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	auth := dtos.Auth{
		Token:        "token",
		ExpiresAt:    time.Now().UnixMilli() + time.Hour.Milliseconds()*24,
		RefreshToken: "refresh token",
	}

	data, err := json.Marshal(auth)

	if err != nil {
		http.Error(w, "Failed to marshal user data", http.StatusInternalServerError)
		return
	}

	w.Write(data)

}
