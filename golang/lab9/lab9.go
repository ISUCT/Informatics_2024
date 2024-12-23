package lab9

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunLab9() {
	reader := bufio.NewReader(os.Stdin)
	var fileName string
	fmt.Print("Введите название файла:\n")
	fmt.Scan(&fileName)
	CreatingFile(fileName)
	taskList := ToDoList{}

	fmt.Println("Список доступных команд:\n" +
		"1. Добавить задачу\n" +
		"2. Показать все задачи\n" +
		"3. Отметить задачу как выполненную\n" +
		"4. Удалить задачу\n" +
		"5. Поиск задачи\n" +
		"6. Выйти")

	for {
		fmt.Print("Введите номер команды\n")
		var command string
		fmt.Scan(&command)
		switch command {
		case "1":
			for {
				fmt.Println("Напишите описание задач (когда закончите напишите 'Нет'):")
				var description string
				description, _ = reader.ReadString('\n')
				description = strings.TrimSpace(description)
				if description == "Нет" {
					fmt.Println("Данные успешно записаны")
					break
				}
				taskList.AddTask(fileName, description)
			}
		case "2":
			taskList.PrintToDoList()
		case "3":
			fmt.Print("Введите номер задачи для обновления статуса:\n")
			var index int
			fmt.Scan(&index)
			taskList.StatusСhange(index)
		case "4":
			fmt.Print("Введите номер задачи, которую хотите удалить:\n")
			var index int
			fmt.Scan(&index)
			taskList.DeletingTask(index)
		case "5":
			fmt.Print("Напишите ключивое слово для поиска задачи\n")
			var keyWord string
			fmt.Scan(&keyWord)
			task, err := taskList.TaskSearch(keyWord)
			if err != nil {
				fmt.Print("Задача не найдена")
				return
			}
			fmt.Print("По ключевому слову найдена задача:", task, "\n")
		case "6":
			fmt.Print("Программа завершина")
			return
		default:
			fmt.Println("Пожалуйста, выберите номер команды")
		}
	}
}
