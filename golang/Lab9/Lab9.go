package Lab9

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string
	Completed   bool
}

type TaskManager interface {
	AddTask(description string)
	ShowTasks()
	CompleteTask(index int)
	DeleteTask(index int)
	SearchTask(keyword string)
	LoadTasks(filename string) error
	SaveTasks(filename string) error
}

type SimpleTaskManager struct {
	tasks []Task
}

func (tm *SimpleTaskManager) AddTask(description string) {
	tm.tasks = append(tm.tasks, Task{Description: description, Completed: false})
}

func (tm *SimpleTaskManager) ShowTasks() {
	if len(tm.tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}
	for i, task := range tm.tasks {
		status := "не выполнена"
		if task.Completed {
			status = "выполнена"
		}
		fmt.Printf("%d: %s [%s]\n", i+1, task.Description, status)
	}
}

func (tm *SimpleTaskManager) CompleteTask(index int) {
	if index < 0 || index >= len(tm.tasks) {
		fmt.Println("Некорректный индекс.")
		return
	}
	tm.tasks[index].Completed = true
}

func (tm *SimpleTaskManager) DeleteTask(index int) {
	if index < 0 || index >= len(tm.tasks) {
		fmt.Println("Некорректный индекс.")
		return
	}
	tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)
}

func (tm *SimpleTaskManager) SearchTask(keyword string) {
	found := false
	for _, task := range tm.tasks {
		if strings.Contains(task.Description, keyword) {
			fmt.Println(task.Description)
			found = true
		}
	}
	if !found {
		fmt.Println("Задачи не найдены.")
	}
}

func (tm *SimpleTaskManager) LoadTasks(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}
		description := parts[0]
		completed := parts[1] == "1"
		tm.tasks = append(tm.tasks, Task{Description: description, Completed: completed})
	}

	return scanner.Err()
}

func (tm *SimpleTaskManager) SaveTasks(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tm.tasks {
		status := "0"
		if task.Completed {
			status = "1"
		}
		_, err := writer.WriteString(fmt.Sprintf("%s|%s\n", task.Description, status))
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func RunLab9() {
	manager := &SimpleTaskManager{}
	const filename = "tasks.txt"

	if err := manager.LoadTasks(filename); err != nil {
		fmt.Printf("Ошибка загрузки задач: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=============================")
		fmt.Println("         МЕНЮ ЗАДАЧ         ")
		fmt.Println("=============================")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск задачи")
		fmt.Println("6. Выйти")
		fmt.Print("Выберите действие (1-6): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Введите описание задачи: ")
			scanner.Scan()
			description := scanner.Text()
			manager.AddTask(description)
		case "2":
			manager.ShowTasks()
		case "3":
			fmt.Print("Введите номер задачи для отметки как выполненной: ")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			manager.CompleteTask(index - 1)
		case "4":
			fmt.Print("Введите номер задачи для удаления: ")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			manager.DeleteTask(index - 1)
		case "5":
			fmt.Print("Введите ключевое слово для поиска: ")
			scanner.Scan()
			keyword := scanner.Text()
			manager.SearchTask(keyword)
		case "6":
			if err := manager.SaveTasks(filename); err != nil {
				fmt.Printf("Ошибка сохранения задач: %v\n", err)
			}
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, попробуйте снова.")
		}
	}
}
