package runtodocreate

import "Go-Todolist/domain_todolist/model/repository"

type Outport interface {
	repository.SaveTodoRepo
}
