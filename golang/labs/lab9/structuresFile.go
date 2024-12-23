package lab9

type ToDoList struct {
	Tasks []ToDoTask `json:"tasks"`
}

type ToDoTask struct {
	Description string `json:"Description"`
	Status      string `json:"Status"`
}
