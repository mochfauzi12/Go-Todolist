package repository

import (
	"Go-Todolist/domain_todolist/model/entity"
	"context"
)

type SaveTodoRepo interface {
	SaveTodo(ctx context.Context, obj *entity.Todo) error
}

type FinAllTodoRepo interface {
	FinAllTodo(ctx context.Context, obj *entity.Todo) (*entity.Todo, error)
}

type FindOneTodoByIDRepo interface {
	FindOneTodoByID(ctx context.Context, todoID string) (*entity.Todo, error)
}
