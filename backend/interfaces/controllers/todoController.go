package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/kikils/golang-todo/domain/model"
	"github.com/kikils/golang-todo/interfaces/database"
	"github.com/kikils/golang-todo/usecase"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(handler database.Sqlhandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				Sqlhandler: handler,
			},
		},
	}
}

func (controller *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todoReceptor struct {
		Title       string `json:"title"`
		Note        string `json:"note"`
		DueDate     string `json:"due_date"`
		IsCompleted bool   `json:"is_completed"`
	}
	if err := json.Unmarshal(b, &todoReceptor); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	formattedTime := parseTime(todoReceptor.DueDate)
	todo := model.Todo{
		Title:       todoReceptor.Title,
		Note:        todoReceptor.Note,
		DueDate:     formattedTime,
		IsCompleted: todoReceptor.IsCompleted,
	}
	id, err := controller.Interactor.Add(todo)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, id)
}

func (controller *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todoReceptor struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Note        string `json:"note"`
		DueDate     string `json:"due_date"`
		IsCompleted bool   `json:"is_completed"`
	}
	if err := json.Unmarshal(b, &todoReceptor); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	formattedTime := parseTime(todoReceptor.DueDate)
	todo := model.Todo{
		ID:          todoReceptor.ID,
		Title:       todoReceptor.Title,
		Note:        todoReceptor.Note,
		DueDate:     formattedTime,
		IsCompleted: todoReceptor.IsCompleted,
	}
	id, err := controller.Interactor.Update(todo)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, id)
}

func (controller *TodoController) FindAll(w http.ResponseWriter, r *http.Request) {
	todoList, err := controller.Interactor.Todos()
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, todoList)
}

func (controller *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = controller.Interactor.Delete(req.ID)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, "Success!")
}

func ResponseOk(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	body := map[string]string{
		"error": message,
	}
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func parseTime(date string) (formattedTime time.Time) {
	formattedTime, _ = time.Parse("20060102", date)
	return
}
