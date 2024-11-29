package lab9

import (
	"log"

	"isuct.ru/informatics2022/todolist/todo"
)

func CompleteLab9() {
	var list todo.Todo

	err := list.Load()
	if err != nil {
		log.Fatal(err)
	}

	list.OutputTodo()

	list.Save()
}
