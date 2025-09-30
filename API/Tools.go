package gotodoapi

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type Todo struct {
	ID   int    `json:"id,omitempty"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

type API struct {
	Todos []Todo `json:"todos"`
}

// SaveTodos saves the todos to a file
func (a *API) SaveTodos() ([]Todo, error) {
	file, err := os.Create("db.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(a.Todos); err != nil {
		return nil, err
	}
	return a.Todos, nil
}

// GetTodos reads the todos from a file
func (a *API) GetTodos() ([]Todo, error) {
	file, err := os.Open("db.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&a.Todos); err != nil {
		if err.Error() == "EOF" {
			return []Todo{}, nil
		}
		return nil, err
	}
	return a.Todos, nil
}

// AddTodo adds a todo to the list
func (a *API) AddTodo(todo Todo) error {
	var err error
	a.Todos, err = a.GetTodos()
	if err != nil {
		return err
	}
	for _, t := range a.Todos {
		if t.ID == todo.ID {
			return errors.New("ID already exists")
		}
	}
	a.Todos = append(a.Todos, todo)
	a.SaveTodos()
	return nil
}

// DeleteTodo deletes a todo from the list
func (a *API) DeleteTodo(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	for i, todo := range a.Todos {
		if todo.ID == idInt {
			a.Todos = append(a.Todos[:i], a.Todos[i+1:]...)
			a.SaveTodos()
			return nil
		}
	}
	return errors.New("ID not found")
}

// GetTodo returns a todo from the list
func (a *API) GetTodo(id int) Todo {
	for _, todo := range a.Todos {
		if todo.ID == id {
			return todo
		}
	}
	return Todo{}
}

// UpdateTodo updates a todo in the list
func (a *API) UpdateTodo(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	for i, todo := range a.Todos {
		if todo.ID == idInt {
			a.Todos[i].Done = !a.Todos[i].Done
			a.SaveTodos()
			return nil
		}
	}

	return errors.New("ID not found")
}
