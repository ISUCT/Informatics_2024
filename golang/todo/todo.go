package todo

import "fmt"

type Task struct {
	name string
}

func NewTask(name string) *Task {
	t := new(Task)
	t.name = name
	return t
}

func Todo() {
	var List map[int]Task
	fmt.Printf("List: %v\n", List)
}
