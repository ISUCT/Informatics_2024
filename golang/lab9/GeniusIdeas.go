package lab9

type GeniusIdeas interface {
	AddTask(description string)
	ShowTasks()
	UpdateTaskStatus(index int)
	DeleteTask(index int)
	SearchTask(word string)
}
