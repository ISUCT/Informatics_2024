package lab9

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filename = "tasks.json"

func CreateFile(filename string) (string, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return "", fmt.Errorf("создание файла: файл уже существует")
	}

	f, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("создание файла: %w", err)
	}
	defer f.Close()

	return filename, nil
}

func ReadTasks() ([]Task, error) {
	tasks, err := ReadFromFileJSON(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("Ошибка чтения задач из файла: %w", err)
	}
	return tasks, nil
}

func WriteTasks(tasks []Task) error {
	err := WriteToFileJSON(filename, tasks)
	if err != nil {
		return fmt.Errorf("Ошибка записи в файл: %w", err)
	}
	return nil
}

func AddTask() {
	var notion string
	fmt.Print("Введите название вашей заметки: ")
	fmt.Scanln(&notion)

	var description string
	fmt.Print("Введите описание вашей заметки: ")
	fmt.Scanln(&description)

	var statusInput string
	var status bool
	for {
		fmt.Print("Введите статус вашей заметки (выполнено/не выполнено): ")
		fmt.Scanln(&statusInput)
		statusInput = strings.ToLower(strings.TrimSpace(statusInput))
		switch statusInput {
		case "выполнено":
			status = true
		case "не выполнено":
			status = false
		default:
			fmt.Println("Неверный ввод. Пожалуйста, введите 'выполнено' или 'не выполнено'.")
			continue
		}
		break
	}

	task := Task{
		Notion:      notion,
		Description: description,
		Status:      status,
	}

	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks = append(tasks, task)
	err = WriteTasks(tasks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Задача успешно добавлена в файл %s\nНазвание заметки: %s\nОписание заметки: %s\nСтатус заметки: %t\n", filename, task.Notion, task.Description, task.Status)
}

func ExitTask() {
	var exitTask string
	for {
		fmt.Print("Если хотите выйти из заметок напишите(да/нет): ")
		fmt.Scanln(&exitTask)
		switch exitTask {
		case "да":
			os.Exit(0)
		case "нет":
			fmt.Println("Вы остались в заметках")
			return
		default:
			fmt.Println("Неверный ввод. Пожалуйста, введите 'да' или 'нет'.")
		}
	}
}

func ShowJSONFile() {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println(string(content))
}

func DeleteTask() {
	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println(err)
		return
	}

	var index int
	fmt.Print("Введите номер задачи для удаления: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(tasks) {
		fmt.Println("Неверный номер задачи")
		return
	}

	tasks = append(tasks[:index-1], tasks[index:]...)
	err = WriteTasks(tasks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Задача успешно удалена")
}

func UpdateTaskStatus() {
	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println(err)
		return
	}

	var index int
	fmt.Print("Введите номер задачи для обновления статуса: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(tasks) {
		fmt.Println("Неверный номер задачи")
		return
	}

	tasks[index-1].Status = true
	err = WriteTasks(tasks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Статус задачи успешно обновлен")
}

func SearchTask() {
	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println(err)
		return
	}

	var searchText string
	fmt.Print("Введите текст для поиска: ")
	searchText, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	searchText = strings.TrimSpace(searchText)

	found := false
	for _, task := range tasks {
		if strings.Contains(task.Notion, searchText) || strings.Contains(task.Description, searchText) {
			var status string
			switch task.Status {
			case true:
				status = "выполнена"
			case false:
				status = "не выполнена"
			}
			fmt.Printf("Название: %s\nОписание: %s\nСтатус: %s\n", task.Notion, task.Description, status)
			found = true
		}
	}

	if !found {
		fmt.Println("Задача не найдена")
	}
}

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
