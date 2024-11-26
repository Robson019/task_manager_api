package request

type TaskDTO struct {
	Title       string `json:"title" example:"Task Title"`
	Description string `json:"description" example:"Task Description"`
	Status      string `json:"status" example:"pending"`
}
