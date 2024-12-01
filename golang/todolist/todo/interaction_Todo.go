package todo

import (
	"fmt"
	"sort"
	"time"
)

func (T *Todo) OutputTodo() {
	for i, task := range T.List {
		fmt.Printf("%v | %v ; осталось %v д. ; тег: %v\n", i, task.Name, task.getDays(), task.Teg)
	}
	fmt.Println()
}

func (T *Task) getDays() int {
	return int((T.Deadline.Sub(T.DateCreate) / time.Hour) / 24)
}

func (T *Todo) SortReverse() {
	var ReverseList []Task

	for i := len(T.List) - 1; i >= 0; i-- {
		ReverseList = append(ReverseList, T.List[i])
	}
	T.List = ReverseList
}

func (T *Todo) SortTaskNameAlphabeticalOrder() {
	sort.Slice(T.List,
		func(i, j int) bool {
			return T.List[i].Name < T.List[j].Name
		})
}

func (T *Todo) SortTaskDataCreate() {
	sort.Slice(T.List,
		func(i, j int) bool {
			return T.List[i].DateCreate.Unix() < T.List[j].DateCreate.Unix()
		})
}

func (T *Todo) SortTaskDeadline() {
	sort.Slice(T.List,
		func(i, j int) bool {
			return T.List[i].Deadline.Unix() > T.List[j].Deadline.Unix()
		})
}
