package database

import (
	"github.com/kikils/golang-todo/domain/model"
)

type TodoRepository struct {
	Sqlhandler
}

func (repo *TodoRepository) Store(t model.Todo) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		"INSERT INTO todos (title, note, duedate, is_completed, user_id) VALUES ($1,$2,$3,$4,$5) RETURNING id;", t.Title, t.Note, t.DueDate, t.IsCompleted, t.UserID,
	)

	if err != nil {
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *TodoRepository) Update(t model.Todo) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		"UPDATE todos SET title=$1, note=$2, duedate=$3, is_completed=$4 WHERE id=$5 RETURNING id;",
		t.Title, t.Note, t.DueDate, t.IsCompleted, t.ID,
	)

	if err != nil {
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *TodoRepository) Delete(id int) (err error) {
	_, err = repo.Sqlhandler.Query("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return
	}
	return
}

func (repo *TodoRepository) FindByUserID(userID string) (todoList model.Todos, err error) {
	rows, err := repo.Sqlhandler.Query("SELECT id, title, note, to_char(duedate, 'YYYY/MM/DD'), is_completed FROM todos WHERE user_id=$1 ORDER BY id;", userID)
	if err != nil {
		return
	}
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&(todo.ID), &(todo.Title), &(todo.Note), &(todo.DueDate), &(todo.IsCompleted)); err != nil {
			continue
		}
		todoList = append(todoList, todo)
	}
	return
}
