package taskutilis

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	structure "isuct.ru/informatics2022/labs/lab9/taskstruct"
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

func AddTask(filename string, description string) error {
	var tasks []structure.Task
	data, err := os.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			return fmt.Errorf("ошибка при разборе JSON: %w", err)
		}
	}
	tasks = append(tasks, structure.Task{Description: description, Status: false})
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("ошибка при создании JSON: %w", err)
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %w", err)
	}
	return nil
}

func DeleteTask(index int, tasks *[]structure.Task) {
	if index < 0 || index >= len(*tasks) {
		fmt.Println("Неверный номер задачи.")
		return
	}
	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
}

func ShowTasks(tasks []structure.Task) {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s (%t)\n", i+1, task.Description, task.Status)
	}
}

func SearchTask(filename string, searchingTask string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == searchingTask {
			found = true
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("ошибка при чтении файла: %w", err)
	}
	if found {
		return "Задача найдена", nil
	} else {
		return "Задача не найдена", nil
	}
}

func UpdateTaskStatus(index int, tasks *[]structure.Task) {
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

func SaveTasks(filename string, tasks []structure.Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Ошибка при сохранении данных:", err)
		return
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}
