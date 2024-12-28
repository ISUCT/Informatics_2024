package lab9

type ToDoList struct {
	Tasks []Exercise `json:"tasks"`
}

type Exercise struct {
	Description string `json:"description"`
	Status      string `json:"status"`
}
