package usecase

import (
	"github.com/kikils/golang-todo/domain/model"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Add(t model.Todo) (id int, err error) {
	id, err = interactor.TodoRepository.Store(t)
	return
}

func (interactor *TodoInteractor) Update(t model.Todo) (id int, err error) {
	id, err = interactor.TodoRepository.Update(t)
	return
}

func (interactor *TodoInteractor) Delete(id int) (err error) {
	err = interactor.TodoRepository.Delete(id)
	return
}

func (interactor *TodoInteractor) FindTodoByUserID(userID string) (todos model.Todos, err error) {
	todos, err = interactor.TodoRepository.FindByUserID(userID)
	return
}
