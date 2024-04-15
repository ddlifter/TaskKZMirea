package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ddlifter/TestForKZMirea/internal/services"

	"github.com/ddlifter/TestForKZMirea/internal/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := services.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// db := config.Connect()
	// defer db.Close()

	w.WriteHeader(http.StatusCreated)
}

//func LoginHandler()
