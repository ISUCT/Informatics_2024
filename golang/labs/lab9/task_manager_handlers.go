package lab9

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

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
