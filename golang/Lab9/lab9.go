package lab9

import (
	"fmt"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

type TaskList struct {
	Tasks []Task
}

func (t *Tasklist) AddTask(description string) {
	id := len(t.Tasks) + 1
	t.Tasks = append(t.Tasks, Task{ID: id, Description: description, Completed: false})
	fmt.Printf("Task added: %d - %s", id, description)
}

func (t *Tasklist) CompleteTask(id int) {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks[i].Completed = true
			fmt.Printf("Task %d completed", id)
			return
		}
	}
	fmt.Printf("Task %d not found", id)
}

func (t *Tasklist) ShowTasks() {
	for _, task := range t.Tasks {
		status := "incomplete"
		if task.Completed {
			status = "completed"
		}
		fmt.Printf("%d: %s [%s]", task.ID, task.Description, status)
	}
}

func RunLab9() {
	taskList := &Tasklist{}
	taskList.AddTask("Write a Go program")
	taskList.AddTask("Test the program")
	taskList.ShowTasks()
	taskList.CompleteTask(1)
	taskList.ShowTasks()
}
