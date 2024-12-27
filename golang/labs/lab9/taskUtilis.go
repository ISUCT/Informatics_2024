package lab9

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func CreateFile(filename string) (string, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return "", fmt.Errorf("файл '%s' уже существует", filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer file.Close()
	return filename, nil
}

func AddTask(tasks *[]Task, description string) {
	*tasks = append(*tasks, Task{Description: description, Status: false})
}

func LoadTasks(filename string, tasks *[]Task) error {
	data, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %w", err)
	}
	err = json.Unmarshal(data, tasks)
	if err != nil {
		return fmt.Errorf("ошибка при разборе JSON: %w", err)
	}
	return nil
}

func DeleteTask(index int, tasks *[]Task) {
	if index < 0 || index >= len(*tasks) {
		fmt.Println("Неверный номер задачи.")
		return
	}
	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
}

func ShowTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s (%t)\n", i+1, task.Description, task.Status)
	}
}

func SearchTask(tasks []Task, keyword string) ([]Task, error) {
	results := []Task{}
	for _, task := range tasks {
		if containsCI(task.Description, keyword) {
			results = append(results, task)
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("задача не найдена")
	}
	return results, nil
}

func containsCI(s, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}

func UpdateTaskStatus(index int, tasks *[]Task) {
	ShowTasks(*tasks)
	index--
	if index < 0 || index >= len(*tasks) {
		fmt.Println("Неверный номер задачи.")
		return
	}
	if (*tasks)[index].Status {
		(*tasks)[index].Status = false
	} else {
		(*tasks)[index].Status = true
	}
}

func SaveTasks(filename string, tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Ошибка при сохранении данных:", err)
		return
	}
	err = os.WriteFile(filename, data, 0600)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func InitTasks(filename string) ([]Task, error) {
	_, err := CreateFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании/проверке файла: %w", err)
	}

	var tasks []Task
	err = LoadTasks(filename, &tasks)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("ошибка при загрузке данных: %w", err)
	}
	return tasks, nil
}
