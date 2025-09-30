package gotodoapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// todosHandler handles the /todos endpoint
func (S *Server) todosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		todos, err := S.api.GetTodos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	} else if r.Method == "POST" {
		todos, err := S.api.GetTodos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		S.api.Todos = todos
		var todo Todo
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		lastID := S.api.Todos[len(S.api.Todos)-1].ID
		todo.ID = lastID + 1
		fmt.Println(lastID, todo.ID)

		err = S.api.AddTodo(todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// todosModifyHandler handles the /todos/:id endpoint
func (S *Server) todosModifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id := r.URL.Path[len("/todos/"):]
		_, err := S.api.GetTodos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = S.api.UpdateTodo(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else if r.Method == "DELETE" {
		id := r.URL.Path[len("/todos/"):]
		_, err := S.api.GetTodos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = S.api.DeleteTodo(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
