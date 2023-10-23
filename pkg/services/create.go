package services

import (
	"crud-go/db"
	"crud-go/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	statement, err := database.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		http.Error(w, "Error creating a statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		http.Error(w, "Error executing the statement", http.StatusInternalServerError)
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Error getting the inserted ID", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User successfully inserted with ID: %d", userID)))
}
