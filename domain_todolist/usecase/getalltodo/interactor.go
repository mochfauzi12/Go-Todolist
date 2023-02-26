package getalltodo

import (
	"context"
	"fmt"
)

type getAllTodoInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getAllTodoInteractor{
		outport: outputPort,
	}
}

func (r *getAllTodoInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	todoObj, err := r.outport.FinAllTodo(ctx, nil)
	if err != nil {
		return nil, err
	}
	if todoObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	return res, nil
}
