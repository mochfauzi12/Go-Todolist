package runtodocheck

import (
	"Go-Todolist/domain_todolist/model/entity"
	"Go-Todolist/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TodoCreateRequest
}

type InportResponse struct {
	Todo *entity.Todo
}
