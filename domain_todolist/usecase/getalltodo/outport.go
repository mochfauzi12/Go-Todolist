package getalltodo

import "Go-Todolist/domain_todolist/model/repository"

type Outport interface {
	repository.FinAllTodoRepo
}
