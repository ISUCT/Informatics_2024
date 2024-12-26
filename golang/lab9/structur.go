package lab9

type ToDoList struct {
	Tasks []Exercise `json:"tasks"`
}

type Exercise struct {
	description string
	status      string
}
