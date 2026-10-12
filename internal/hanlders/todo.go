package hanlders

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dubey-Ji/todoapp/internal/todo"
)

var todos = []todo.Todo{}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// return h.todoService.Create(ctx, todo)
	var newTodo todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "invalid todo body", http.StatusBadRequest)
		return
	}
	todos = append(todos, newTodo)
	w.Header().Set("Content-type", "application/json")
	response := struct {
		Status string    `json:"status"`
		Data   todo.Todo `json:"data"`
	}{
		Status: "ok",
		Data:   newTodo,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}

	// return todo, nil
}

// func UpdateTodo(ctx context.Context, id int64) (todo.Todo, error) {
// 	// return h.todoService.Update(ctx, id)

// }

func ListTodos(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status string      `json:"status"`
		Data   []todo.Todo `json:"data"`
	}{
		Status: "ok",
		Data:   todos,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "could not encode", http.StatusInternalServerError)
		return
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// return h.todoService.Delete(ctx, id)
	var idStr string
	idStr = r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	fmt.Println("deleting todo for id", id)
	for i, v := range todos {
		if v.ID == int(id) {
			todos = append(todos[:i], todos[i+i:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Todo with ID %d deleted successfully", id)
			return
		}
	}
	http.Error(w, "Ivalid ID", http.StatusBadRequest)
}
