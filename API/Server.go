package gotodoapi

import (
	"net/http"
)

type Server struct {
	api API
}

// Run starts the server
func (S *Server) Run(port string) {
	S.api = API{}
	http.HandleFunc("/todos", S.todosHandler)
	http.HandleFunc("/todos/", S.todosModifyHandler)
	http.ListenAndServe(":"+port, nil)
}
