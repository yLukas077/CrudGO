package services

import (
	"crud-go/db"
	"encoding/json"
	"net/http"
	"strconv"
	"crud-go/pkg/models"

	"github.com/gorilla/mux"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	rows, err := database.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, "Error scanning user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error converting users to JSON", http.StatusInternalServerError)
		return
	}
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		http.Error(w, "Error converting parameter to integer", http.StatusBadRequest)
		return
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	row := database.QueryRow("SELECT * FROM users WHERE id = ?", ID)
	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		http.Error(w, "Error scanning user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
		return
	}
}
