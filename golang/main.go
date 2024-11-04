package main

import (
	"isuct.ru/informatics2022/todo"
	"isuct.ru/informatics2022/todo/console"
)

func main() {
	var list todo.Todo

	list.AddTask(console.InputTask())
	list.AddTask(console.InputTask())

	list.OutputTodo()
}
