package lab9

import (
	"encoding/json"
	"fmt"
	"os"

	structure "isuct.ru/informatics2022/labs/lab9/taskstruct"
	"isuct.ru/informatics2022/labs/lab9/taskutilis"
)

func RunLab9() {
	var filename string
	fmt.Print("Введите название файла: ")
	fmt.Scan(&filename)
	taskutilis.CreateFile(filename)
	var tasks []structure.Task
	if _, err := os.Stat(filename); err == nil {
		data, err := os.ReadFile(filename)
		if err == nil {
			err = json.Unmarshal(data, &tasks)
			if err != nil {
				fmt.Println("Ошибка при загрузке данных:", err)
			}
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
				var filename string
				fmt.Print("Введите название файла: ")
				fmt.Scan(&filename)
				taskutilis.CreateFile(filename)
				var description string
				fmt.Print("Введите описание задачи: ")
				fmt.Scan(&description)
				taskutilis.AddTask(filename, description)
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
				var searchingTask string
				fmt.Print("Введите ключевое слово для поиска: ")
				fmt.Scan(&searchingTask)
				taskutilis.SearchTask(filename, searchingTask)
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
