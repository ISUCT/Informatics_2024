package Lab9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunLab9() {
	taskManager := TaskManager{}
	taskManager.LoadFromFile("tasks.json")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск задачи")
		fmt.Println("6. Выйти")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Введите описание задачи:")
			scanner.Scan()
			description := scanner.Text()
			taskManager.AddTask(description)
		case "2":
			taskManager.ShowTasks()
		case "3":
			fmt.Println("Введите номер задачи для отметки как выполненной:")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err == nil {
				taskManager.CompleteTask(index - 1) // Индексация с 0
			} else {
				fmt.Println("Ошибка: некорректный ввод.")
			}
		case "4":
			fmt.Println("Введите номер задачи для удаления:")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err == nil {
				taskManager.DeleteTask(index - 1) // Индексация с 0
			} else {
				fmt.Println("Ошибка: некорректный ввод.")
			}
		case "5":
			fmt.Println("Введите ключевое слово для поиска:")
			scanner.Scan()
			keyword := scanner.Text()
			taskManager.SearchTask(keyword)
		case "6":
			taskManager.SaveToFile("tasks.json")
			fmt.Println("Задачи сохранены. Выход из программы.")
			return
		default:
			fmt.Println("Ошибка: некорректный выбор.")
		}
	}
}
