package lab9

import (
	"fmt"

	"isuct.ru/informatics2022/labs/lab9/TaskUtilis"
)

func RunLab9() {
	var filename string
	fmt.Print("Введите название файла: ")
	fmt.Scanln(&filename)

	tasks, err := TaskUtilis.InitTasks(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск задачи")
		fmt.Println("6. Выйти")

		var choice int
		fmt.Print("Выберите пункт меню: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var description string
			fmt.Print("Введите описание задачи: ")
			fmt.Scan(&description)
			TaskUtilis.AddTask(&tasks, description)
		case 2:
			TaskUtilis.ShowTasks(tasks)
		case 3:
			var index int
			fmt.Print("Введите номер задачи: ")
			fmt.Scan(&index)
			TaskUtilis.UpdateTaskStatus(index, &tasks)

		case 4:
			var index int
			fmt.Print("Введите номер задачи: ")
			fmt.Scan(&index)
			TaskUtilis.DeleteTask(index, &tasks)
		case 5:
			var keyword string
			fmt.Print("Введите ключевое слово для поиска: ")
			fmt.Scan(&keyword)
			TaskUtilis.SearchTask(tasks, keyword)
		case 6:
			TaskUtilis.SaveTasks(filename, tasks)
			fmt.Println("Выход из программы...")
			return
		default:
			fmt.Println("Неверный выбор пункта меню.")
		}
	}
}
