package services

import (
	"crud-go/db"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"crud-go/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		http.Error(w, "Error converting parameter to integer", http.StatusBadRequest)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read the request body", http.StatusInternalServerError)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		http.Error(w, "Error converting user to struct", http.StatusBadRequest)
		return
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	statement, err := database.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Error creating a statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		http.Error(w, "Error executing a statement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
