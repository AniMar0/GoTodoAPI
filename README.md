# GoTodoAPI

## Overview

GoTodoAPI is a simple RESTful API built with Go for managing a Todo list.  
It supports full CRUD operations: creating, reading, updating, and deleting todos, with data stored in a JSON file (`db.json`).

## Features

- **GET /todos**: Returns all todos
- **POST /todos**: Adds a new todo
- **PUT /todos/{id}**: Toggles a todo's done status
- **DELETE /todos/{id}**: Deletes a todo by ID
- Data is stored in a simple JSON file
- Lightweight, easy to extend, and beginner-friendly

## Tech Stack

- Language: Go
- Storage: JSON file (`db.json`)
- HTTP: `net/http` package

## Installation

1. Clone the repository:

```bash
git clone https://github.com/AniMar0/GoTodoAPI.git
```

2. Navigate to the project folder:

```bash
cd GoTodoAPI
```

3. Run the server:

```bash
go run main.go
```

The API will be available at: http://localhost:8080
You can change the port in `main.go` if needed.

## Usage

GET all todos

```bash
curl http://localhost:8080/todos
```

POST a new todo

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"task":"Learn Go","done":false}'
```

PUT update todo (toggle done)

```bash
curl -X PUT http://localhost:8080/todos/1
```

DELETE a todo

```bash
curl -X DELETE http://localhost:8080/todos/1
```

## Future Improvements

- Switch to a real DBMS (SQLite / PostgreSQL)
- Add authentication and authorization
- Make operations concurrency-safe
- Add unit tests for handlers and functions

## Author

Zakaria Kahlaoui
