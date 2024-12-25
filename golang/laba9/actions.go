package laba9

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func (t *TaskList) AddTask(description string) {
	t.Tasks = append(t.Tasks, Task{Description: description, IsDone: false})
}

func LoadData(filename string) (*TaskList, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks TaskList
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}
	return &tasks, nil
}

func (t *TaskList) UpdateTaskStatus(index int) {
	if index < 0 || index >= len(t.Tasks) {
		fmt.Println("Задача с таким индексом не найдена.")
		return
	}
	t.Tasks[index].IsDone = true
	fmt.Println("Статус задачи обновлен на 'Выполнена'.")
}

func (t *TaskList) SearchTask(word string) {
	found := false
	for i, task := range t.Tasks {
		if strings.Contains(strings.ToLower(task.Description), strings.ToLower(word)) {
			status := "Не выполнена"
			if task.IsDone {
				status = "Выполнена"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
			found = true
		}
	}
	if !found {
		fmt.Println("Задача не найдена.")
	}
}

func (t *TaskList) ShowTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("Нет задач.")
		return
	}
	for i, task := range t.Tasks {
		status := "Не выполнена"
		if task.IsDone {
			status = "Выполнена"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
	}
}

func SaveData(filename string, tasks *TaskList) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}

func (t *TaskList) DeleteTask(index int) {
	if index < 0 || index >= len(t.Tasks) {
		fmt.Println("Задача с таким индексом не найдена.")
		return
	}
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
	fmt.Println("Задача удалена.")
}
