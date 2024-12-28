package lab9

type task struct {
	ID          int
	Description string
	Completed   bool
}

func NewTask(id int, description string) Task {
	return Task{
		ID:          id,
		Description: description,
		Completed:   false,
	}
}

func (t *Task) MarkAsComplete() {
	t.Completed = true
}
