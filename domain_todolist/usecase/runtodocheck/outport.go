package runtodocheck

import "Go-Todolist/domain_todolist/model/repository"

type Outport interface {
	repository.FindOneTodoByIDRepo
	repository.SaveTodoRepo
}
