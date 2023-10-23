package handlers

import (
	"crud-go/pkg/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/users", services.CreateUser).Methods("POST")
	router.HandleFunc("/users", services.SearchUsers).Methods("GET")
	router.HandleFunc("/users/{id}", services.SearchUser).Methods("GET")
	router.HandleFunc("/users/{id}", services.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", services.DeleteUser).Methods("DELETE")

	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
