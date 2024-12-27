package lab9

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskManager struct {
	Tasks []Task
}

func handleTaskAction(scanner *bufio.Scanner, prompt string, action func(int)) {
	fmt.Print(prompt)
	if scanner.Scan() {
		var index int
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &index); err == nil {
			action(index)
		} else {
			fmt.Println("Некорректный ввод. Введите целое число.")
		}
	}
}

func RunLab9Task() {
	const filename = "labs/lab9/tasks.json"

	manager := TaskManager{}

	err := manager.LoadTasks(filename)

	if err != nil {
		fmt.Println("Ошибка при загрузке задач:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск задачи")
		fmt.Println("6. Выйти")
		fmt.Print("Выберите действие: ")

		if !scanner.Scan() {
			break
		}

		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Введите описание задачи: ")
			if scanner.Scan() {
				manager.AddTask(scanner.Text())
			}
		case "2":
			manager.ShowTasks()
		case "3":
			handleTaskAction(scanner, "Введите номер задачи для отметки как выполненной: ", func(index int) {
				manager.MarkTaskAsCompleted(index - 1)
			})

		case "4":
			handleTaskAction(scanner, "Введите номер задачи для удаления: ", func(index int) {
				manager.DeleteTask(index - 1)
			})
		case "5":
			fmt.Print("Введите ключевое слово")
			if scanner.Scan() {
				manager.SearchTask(scanner.Text())
			}
		case "6":
			err := manager.SaveTasks(filename)

			if err != nil {
				fmt.Println("Ошибка при сохранении задач:", err)
			}
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте снова.")
		}
	}
}
