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

func (T *Todo) OutputTodo() []Task {
	for i, task := range T.list {
		fmt.Printf("%v %v\n", i, task)
	}
	return T.list
}

func (T *Todo) AddTask(newName string, newTime int, newTeg string) {
	T.list = append(
		[]Task{
			{
				name: newName,
				time: newTime,
				teg:  newTeg,
			},
		},
		T.list...,
	)
}

func (T *Todo) RemoveTask(num int) {
	T.list = append(T.list[:num], T.list[(num+1):]...)
}

func (T *Todo) MoveTask(from int, to int) {
	list := make([]Task, from)
	copy(list, T.list[:from])

	list = append(list, T.list[(from+1):to]...)
	list = append(list, T.list[from])
	list = append(list, T.list[to:]...)

	T.list = list
}

func (T *Todo) EditTask(num int, newName string, newTime int, newTeg string) {
	T.list[num].name = newName
	T.list[num].time = newTime
	T.list[num].teg = newTeg
}

func (T *Todo) EditName(num int, newName string) {
	T.list[num].name = newName
}

func (T *Todo) EditTime(num int, newTime int) {
	T.list[num].time = newTime
}

func (T *Todo) EditTeg(num int, newTeg string) {
	T.list[num].teg = newTeg
}
