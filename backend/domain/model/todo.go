package model

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
	DueDate     string `json:"due_date"`
	IsCompleted bool   `json:"is_completed"`
}

type Todos []Todo
