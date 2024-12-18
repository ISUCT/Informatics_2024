package lab9

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteToFileJSON(filename string, tasks []Task) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Ошибка открытия файла для записи: %w", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return fmt.Errorf("Ошибка при кодировании JSON: %w", err)
	}
	return nil
}

func ReadFromFileJSON(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("Ошибка открытия файла для чтения: %w", err)
	}
	defer file.Close()
	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при декодировании JSON: %w", err)
	}
	return tasks, nil
}
