package lab9

import "fmt"

type Tasklist struct {
	Tasks []Task
}

func (tl *Tasklist) Add(description string) {
	id := len(tl.Tasks) + 1
	tl.Tasks = append(tl.Tasks, NewTask(id, description))
	fmt.Printf("Task added: %d - %s", id, description)
}

func (tl *Tasklist) Complete(id int) {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].MarkAsComplete()
			fmt.Printf("Task %d marked as complete", id)
			return
		}
	}
	fmt.Printf("Task %d not found", id)
}

func (tl *Tasklist) Display() {
	for _, task := range tl.Tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Complete"
		}
		fmt.Printf("%d: %s [%s]", task.ID, task.Description, status)
	}
}
