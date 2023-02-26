package withgorm

import (
	"Go-Todolist/domain_todolist/model/entity"
	"Go-Todolist/shared/config"
	"Go-Todolist/shared/gogen"
	"Go-Todolist/shared/infrastructure/logger"
	"context"
)

type gateway struct {
	appData gogen.ApplicationData
	config  *config.Config
	log     logger.Logger
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) FinAllTodo(ctx context.Context, obj *entity.Todo) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) FindOneTodoByID(ctx context.Context, todoID string) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	return nil
}
