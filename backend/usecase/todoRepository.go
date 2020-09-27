package usecase

import (
	"github.com/kikils/golang-todo/domain/model"
)

type TodoRepository interface {
	Store(model.Todo) (int, error)
	Update(t model.Todo) (id int, err error)
	Delete(id int) (err error)
	FindAll() (model.Todos, error)
}
