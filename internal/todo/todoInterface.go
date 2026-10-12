package todo

import "context"

type Todo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TodoInterface interface {
	Create(ctx context.Context, todo Todo) (Todo, error)
	Update(ctx context.Context, id int64) (Todo, error)
	List(ctx context.Context) ([]Todo, error)
	Delete(ctx context.Context, id int64) error
}
