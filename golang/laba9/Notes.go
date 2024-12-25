package laba9

type Notes interface {
	AddTask(description string)
	ShowTasks()
	UpdateTaskStatus(index int)
	DeleteTask(index int)
	SearchTask(word string)
}
