package lab9

import (
	"isuct.ru/informatics2022/todolist/todo"
)

func CompleteLab9() {
	var list todo.Todo

	list.AddTask("0", 0, "0тег")
	list.AddTask("1", 1, "1тег")
	list.AddTask("2", 2, "2тег")

	list.OutputTodo()
}
