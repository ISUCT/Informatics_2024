package lab9

import (
	"fmt"
	"os"

	"isuct.ru/informatics2022/labs/lab9/taskstruct"
	"isuct.ru/informatics2022/labs/lab9/taskutilis"
)

func RunLab9() {
	var filename string
	fmt.Print("Введите название файла: ")
	fmt.Scanln(&filename)
	_, err := taskutilis.CreateFile(filename)
	if err != nil {
		fmt.Println("Ошибка при создании/проверке файла:", err)
		return
	}
	var tasks []taskstruct.Task
	err = taskutilis.LoadTasks(filename, &tasks)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Ошибка при загрузке данных:", err)

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
				taskutilis.AddTask(&tasks, description)
			case 2:
				taskutilis.ShowTasks(tasks)
			case 3:
				var index int
				fmt.Print("Введите номер задачи: ")
				fmt.Scan(&index)
				taskutilis.UpdateTaskStatus(index, &tasks)

			case 4:
				var index int
				fmt.Print("Введите номер задачи: ")
				fmt.Scan(&index)
				taskutilis.DeleteTask(index, &tasks)
			case 5:
				var keyword string
				fmt.Print("Введите ключевое слово для поиска: ")
				fmt.Scan(&keyword)
				taskutilis.SearchTask(tasks, keyword)
			case 6:
				taskutilis.SaveTasks(filename, tasks)
				fmt.Println("Выход из программы...")
				return
			default:
				fmt.Println("Неверный выбор пункта меню.")
			}
		}
	}
}
