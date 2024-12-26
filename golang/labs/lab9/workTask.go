package lab9

import (
	"fmt"
	"os"
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
