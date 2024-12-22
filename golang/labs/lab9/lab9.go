package lab9

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskManager struct {
	Tasks []Task
}

func (tm *TaskManager) AddTask(description string) {
	tm.Tasks = append(tm.Tasks, Task{Description: description, Completed: false})
}

func (tm *TaskManager) ShowTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}

	for i, task := range tm.Tasks {
		status := "Не завершена"

		if task.Completed {
			status = "Завершена"
		}

		fmt.Printf("%d. %s %s\n", i+1, task.Description, status)
	}
}

func (tm *TaskManager) MarkTaskAsCompleted(index int) {
	if index < 0 || index >= len(tm.Tasks) {
		fmt.Println("Некорректный индекс для массива")
		return
	}

	tm.Tasks[index].Completed = true
	fmt.Println("Задача выполнена, как выполненная")
}

func (tm *TaskManager) DeleteTask(index int) {
	if index < 0 || index >= len(tm.Tasks) {
		fmt.Println("Некорректный индекс для массива")
		return
	}

	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]...)
	fmt.Println("Задача удалена")
}

func (tm *TaskManager) SearchTask(keyword string) {
	found := false
	for i, task := range tm.Tasks {
		if strings.Contains(strings.ToLower(task.Description), strings.ToLower((keyword))) {
			status := "Не выполнена"
			if task.Completed {
				status = "Выполнена"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
			found = true
		}
	}

	if !found {
		fmt.Println("Задачи не найдены")
	}
}

func (tm *TaskManager) LoadTasks(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&tm.Tasks)
}

func (tm *TaskManager) SaveTasks(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tm.Tasks)
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
			fmt.Print("Введите номер задачи: ")
			if scanner.Scan() {
				var index int
				_, err := fmt.Sscanf(scanner.Text(), "%d", &index)
				if err == nil {
					manager.MarkTaskAsCompleted(index - 1)
				} else {
					fmt.Print("Введите целое число")
				}
			}
		case "4":
			fmt.Print("Введите номер задачи: ")
			if scanner.Scan() {
				var index int
				_, err := fmt.Sscanf(scanner.Text(), "%d", &index)
				if err == nil {
					manager.DeleteTask(index - 1)
				} else {
					fmt.Print("Введите целое число")
				}
			}
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
