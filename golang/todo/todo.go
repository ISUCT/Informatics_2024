package todo

import (
	"fmt"
)

type Task struct {
	name string
	time int
	teg  string
}

type Todo struct {
	list []Task
}

func (T *Todo) OutputTodo() {
	for i, task := range T.list {
		fmt.Printf("%v %v\n", i+1, task)
	}
}

func (T *Todo) AddTask(newName string, newTime int, newTeg string) {
	T.list = append(
		T.list,
		Task{
			name: newName,
			time: newTime,
			teg:  newTeg,
		},
	)
}

func (T *Todo) RemoveTask() {
}
