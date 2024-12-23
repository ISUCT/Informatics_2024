package lab9

type ToDoList struct {
	Tasks []ToDoTask `json:"tasks"`
}

type ToDoTask struct {
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (t *ToDoTask) GetDescription() string {
	return t.Description
}

func (t *ToDoTask) GetStatus() string {
	return t.Status
}

func (t *ToDoTask) SetDescription(description string) {
	t.Description = description
}

func (t *ToDoTask) SetStatus(status string) {
	t.Status = status
}
