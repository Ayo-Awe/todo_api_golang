package dto

type CreateTodoDTO struct {
	Title       string `json:"title" validator:"required,string"`
	Description string `json:"description" validator:"string"`
}

type UpdateTodoDTO struct {
	Title       string `json:"title" validator:"string"`
	Description string `json:"description" validator:"string"`
}

type UpdateTodoStatusDTO struct {
	Completed bool `json:"completed" validator:"required,bool"`
}
