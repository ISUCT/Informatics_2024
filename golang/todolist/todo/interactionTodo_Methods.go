package todo

import "fmt"

func (T *Todo) OutputTodo() []Task {
	for i, task := range T.List {
		fmt.Printf("%v %v\n", i, task)
	}
	return T.List
}

func (T *Todo) SortedTask() {

}
