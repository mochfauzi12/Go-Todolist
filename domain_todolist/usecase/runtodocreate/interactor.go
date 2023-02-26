package runtodocreate

import (
	"Go-Todolist/domain_todolist/model/entity"
	"context"
)

type runTodoCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoCreateInteractor{
		outport: outputPort,
	}
}

func (r *runTodoCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := entity.NewTodo(req.TodoCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	res.Todo = todoObj

	return res, nil
}
