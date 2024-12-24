package TaskStruct

type Task struct {
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func (t *Task) GetStatus() bool {
	return t.Status
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) SetStatus(status bool) {
	t.Status = status
}

func (t *Task) SetDescription(description string) {
	t.Description = description
}
