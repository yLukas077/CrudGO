package services

import (
	"crud-go/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	statement, err := database.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		http.Error(w, "Error creating a statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		http.Error(w, "Error executing a statement", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User deleted successfully!"))
}
