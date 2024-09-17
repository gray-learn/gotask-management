package main

import (
	"log"
	"net/http"

	"gotask-management/internal/database" // {{ edit_1 }}
	"gotask-management/internal/handlers" // {{ edit_2 }}

	"github.com/gorilla/mux"
	// "handlers"
)

func main() {
	// Initialize database
	db, err := database.InitDB("tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create router
	r := mux.NewRouter()

	// Set up routes
	// r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/tasks", handlers.ListTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/create", handlers.CreateTaskHandler).Methods("GET", "POST")
	r.HandleFunc("/tasks/{id}/edit", handlers.EditTaskHandler).Methods("GET", "POST")
	r.HandleFunc("/tasks/{id}/delete", handlers.DeleteTaskHandler).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
	})

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
