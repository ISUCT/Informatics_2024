package lab9

import (
	"fmt"
)

const path string = "lab9/list.json"
const menu string = "\nПожалуйста, введите номер операции, чтобы её выполнить\n1. Добавить задачу\n2. Показать все задачи\n3. Отметить задачу как выполненную\n4. Удалить задачу\n5. Поиск задачи\n6. Выйти"

func ShowLab9() {
	CreateFile(path)
	var choice int
	TaskList := ToDoList{}
	var description string

	for {
		fmt.Println(menu)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Введите описание задачи: ")
			fmt.Scan(&description)
			TaskList.AddExercise(description, path)
			fmt.Println("Задача добавлена")
		case 2:
			TaskList.ShowTodoList()
		case 3:
			fmt.Println("Введите номер выполненной задачи: ")
			var number int
			fmt.Scan(&number)
			TaskList.StatusComplete(number)
			fmt.Printf("Задача %v обновлена\n", number)
		case 4:
			fmt.Println("Введите номер задачи, которую хотите удалить: ")
			var number int
			fmt.Scan(&number)
			TaskList.DeleteExercise(number)
			fmt.Printf("Задача %v удалена\n", number)
		case 5:
			fmt.Println("Введите ключевое слово для поиска задачи: ")
			var word string
			fmt.Scan(&word)
			exerciseFinded := TaskList.SearchExercise(word)
			fmt.Println("Найденная задача: ", exerciseFinded)
		case 6:
			return
		default:
			fmt.Println("Введена неверная команда!")
		}
	}
}
