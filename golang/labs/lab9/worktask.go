package lab9

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
		if statusInput == "выполнено" {
			status = true
			break
		} else if statusInput == "не выполнено" {
			status = false
			break
		} else {
			fmt.Println("Неверный ввод. Пожалуйста, введите 'выполнено' или 'не выполнено'.")
		}
	}

	task := Task{
		Notion:      notion,
		Description: description,
		Status:      status,
	}

	filename := "task.json"
	tasks, err := ReadFromFileJSON(filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Ошибка чтения задач из файла:", err)
		return
	}

	tasks = append(tasks, task)
	err = WriteToFileJSON(filename, tasks)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
	fmt.Printf("Задача успешно добавлена в файл %s\nНазвание заметки: %s\nОписание заметки: %s\nСтатус заметки: %t\n", filename, task.Notion, task.Description, task.Status)
}

func ExitTask() {
	var exitTask string
	for {
		fmt.Print("Если хотите выйти из заметок напишите(да/нет): ")
		fmt.Scanln(&exitTask)
		if exitTask == "да" {
			os.Exit(0)
		} else if exitTask == "нет" {
			fmt.Println("Вы остались в заметках")
			break
		} else {
			fmt.Println("Неверный ввод. Пожалуйста, введите 'да' или 'нет'.")
		}
	}
}

func ShowJSONFile() {
	filename := "task.json"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println(string(content))
}

func DeleteTask() {
	filename := "task.json"
	tasks, err := ReadFromFileJSON(filename)
	if err != nil {
		fmt.Println("Ошибка чтения задач из файла:", err)
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
	err = WriteToFileJSON(filename, tasks)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
	fmt.Println("Задача успешно удалена")
}

func UpdateTaskStatus() {
	filename := "task.json"
	tasks, err := ReadFromFileJSON(filename)
	if err != nil {
		fmt.Println("Ошибка чтения задач из файла:", err)
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
	err = WriteToFileJSON(filename, tasks)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
	fmt.Println("Статус задачи успешно обновлен")
}

func SearchTask() {
	filename := "task.json"
	tasks, err := ReadFromFileJSON(filename)
	if err != nil {
		fmt.Println("Ошибка чтения задач из файла:", err)
		return
	}

	var searchText string
	fmt.Print("Введите текст для поиска: ")
	searchText, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	searchText = strings.TrimSpace(searchText)

	found := false
	for _, task := range tasks {
		if strings.Contains(task.Notion, searchText) || strings.Contains(task.Description, searchText) {
			status := "не выполнена"
			if task.Status {
				status = "выполнена"
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
