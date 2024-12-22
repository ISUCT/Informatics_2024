package lab9

type Genius ideas interface {
	AddTask(description string)
	ShowTasks()
	UpdateTaskStatus(index int)
	DeleteTask(index int)
	SearchTask(word string)
}
