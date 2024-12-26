package lab9

import "fmt"

func ChooseTypeNotion() {
	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск задачи")
		fmt.Println("6. Выйти")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			AddTask()
		case 2:
			ShowJSONFile()
		case 3:
			UpdateTaskStatus()
		case 4:
			DeleteTask()
		case 5:
			SearchTask()
		case 6:
			ExitTask()
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите действие от 1 до 6.")
		}
	}
}
