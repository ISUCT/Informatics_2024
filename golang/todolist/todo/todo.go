package todo

import (
	"fmt"
)

type Task struct {
	Name string `json:"Name"`
	Date int    `json:"Date"`
	Time int    `json:"Time"`
	Teg  string `json:"Teg"`
}

type Todo struct {
	List []Task `json:"list"`
}

func (T *Todo) OutputTodo() []Task {
	for i, task := range T.List {
		fmt.Printf("%v %v\n", i, task)
	}
	return T.List
}

func (T *Todo) AddTask(newName string, newTime int, newTeg string) {
	T.List = append(
		[]Task{
			{
				Name: newName,
				Date: 0,
				Time: newTime,
				Teg:  newTeg,
			},
		},
		T.List...,
	)
}

func (T *Todo) RemoveTask(num int) {
	T.List = append(T.List[:num], T.List[(num+1):]...)
}

func (T *Todo) MoveTask(from int, to int) {
	list := make([]Task, from)
	copy(list, T.List[:from])

	list = append(list, T.List[(from+1):to]...)
	list = append(list, T.List[from])
	list = append(list, T.List[to:]...)

	T.List = list
}

func (T *Todo) EditTask(num int, newName string, newTime int, newTeg string) {
	T.List[num].Name = newName
	T.List[num].Time = newTime
	T.List[num].Teg = newTeg
}

func (T *Todo) EditName(num int, newName string) {
	T.List[num].Name = newName
}

func (T *Todo) EditTime(num int, newTime int) {
	T.List[num].Time = newTime
}

func (T *Todo) EditTeg(num int, newTeg string) {
	T.List[num].Teg = newTeg
}
