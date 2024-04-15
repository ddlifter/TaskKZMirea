package main

import (
	"net/http"

	"github.com/ddlifter/TestForKZMirea/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	//router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	http.ListenAndServe(":8080", nil)
}
