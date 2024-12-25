package Lab9

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type TaskManager struct {
	Tasks []Task
}

type TaskManagerInterface interface {
	AddTask(description string)
	ShowTasks()
	CompleteTask(index int)
	DeleteTask(index int)
	SearchTask(keyword string)
	SaveToFile(filename string)
	LoadFromFile(filename string)
}

func (tm *TaskManager) AddTask(description string) {
	tm.Tasks = append(tm.Tasks, Task{Description: description, Completed: false})
}

func (tm *TaskManager) ShowTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("Нет задач для отображения.")
		return
	}
	for i, task := range tm.Tasks {
		status := "Не выполнена"
		if task.Completed {
			status = "Выполнена"
		}
		fmt.Printf("%d: %s [%s]\n", i+1, task.Description, status)
	}
}

func (tm *TaskManager) CompleteTask(index int) {
	if index < 0 || index >= len(tm.Tasks) {
		fmt.Println("Ошибка: индекс задачи вне диапазона.")
		return
	}
	tm.Tasks[index].Completed = true
}

func (tm *TaskManager) DeleteTask(index int) {
	if index < 0 || index >= len(tm.Tasks) {
		fmt.Println("Ошибка: индекс задачи вне диапазона.")
		return
	}
	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]...)
}

func (tm *TaskManager) SearchTask(keyword string) {
	found := false
	for _, task := range tm.Tasks {
		if strings.Contains(task.Description, keyword) {
			status := "Не выполнена"
			if task.Completed {
				status = "Выполнена"
			}
			fmt.Printf("Найдена задача: %s [%s]\n", task.Description, status)
			found = true
		}
	}
	if !found {
		fmt.Println("Задачи с таким ключевым словом не найдены.")
	}
}

func (tm *TaskManager) SaveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Ошибка при сохранении файла: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tm.Tasks); err != nil {
		fmt.Printf("Ошибка при записи в файл: %v\n", err)
	}
}

func (tm *TaskManager) LoadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tm.Tasks); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}
}
