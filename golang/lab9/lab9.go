package lab9

import "fmt"

func ShowLab9() {
	fmt.Println("Введите название файла")
	var name string
	fmt.Scan(&name)
	path := "lab9/" + name + ".json"

	CreateFile(path)
	var choice int
	TaskList := ToDoList{}
	var description string

	for {
		fmt.Println("\nПожалуйста, введите номер операции, чтобы её выполнить")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск задачи")
		fmt.Println("6. Выйти")
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

