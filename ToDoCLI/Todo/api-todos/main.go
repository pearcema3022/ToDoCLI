package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic for adding a new TODO item
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic for listing all TODO items
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic for deleting a TODO item
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	todoID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == todoID {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func main() {
	// Registering the handler functions for specific URL paths
	http.HandleFunc("/todos", addTodoHandler)
	http.HandleFunc("/todos", listTodosHandler)
	http.HandleFunc("/todos", deleteTodoHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Todo API server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
