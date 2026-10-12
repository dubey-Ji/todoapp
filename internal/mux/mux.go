package mux

import (
	"net/http"

	"github.com/dubey-Ji/todoapp/internal/hanlders"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", hanlders.Healthz)
	mux.HandleFunc("POST /todo", hanlders.CreateTodo)
	mux.HandleFunc("GET /todo", hanlders.ListTodos)
	mux.HandleFunc("DELETE /todo/{id}", hanlders.DeleteTodo)
	// mux.HandleFunc("PUT /todo/:id", hanlders.UpdateTodo)

	return mux
}
