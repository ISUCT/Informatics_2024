package laba9

import (
	"fmt"
)

func Readylab9() {
	const filename = "tasks.json"

	tasks, err := LoadData(filename)
	if err != nil {
		fmt.Println("Ошибка при загрузке данных:", err)
		return
	}

	var taskMenu Notes = tasks

	for {
		fmt.Println("\nМеню доступных задач:")
		fmt.Println("1. Добавить новую задачу")
		fmt.Println("2. Показать все имеющиеся задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск нужной задачи")
		fmt.Println("6. Выйти из программы")
		fmt.Print("Выберите нужное вам действие:")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var zadacha string
			fmt.Print("Введите задачу:")
			fmt.Scan(&zadacha)
			taskMenu.AddTask(zadacha)

		case 2:
			taskMenu.ShowTasks()

		case 3:
			var index int
			fmt.Print("Введите индекс задачи для отметки как выполненной: ")
			fmt.Scan(&index)
			taskMenu.UpdateTaskStatus(index - 1)

		case 4:
			var index int
			fmt.Print("Введите индекс задачи для удаления: ")
			fmt.Scan(&index)
			taskMenu.DeleteTask(index - 1)

		case 5:
			var word string
			fmt.Print("Введите нужное слово для поиска: ")
			fmt.Scan(&word)
			taskMenu.SearchTask(word)

		case 6:
			err := SaveData(filename, tasks)
			if err == nil {
				fmt.Println("Данные успешно сохранены.")
				break
			}
			
			fmt.Println("Не удалось сохранить данные, ошибка:", err)
			return

		default:
			fmt.Println("Выбор не корректен")
		}
	}
}
