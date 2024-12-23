package lab9

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func CreateFile(filename string) (string, error) {
	_, err := os.Stat(filePath + filename)
	if err == nil {
		return filename, nil
	}
	f, err := os.Create(filePath + filename)
	if err != nil {
		return "", fmt.Errorf("создание файла: %w", err)
	}
	defer f.Close()
	return filename, nil
}

func ReadDataFromFile(filename string, list *ToDoList) error {
	var readedList ToDoList
	f, err := os.Open(filePath + filename)
	if err != nil {
		return fmt.Errorf("открытие файла: %w", err)
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		err = json.Unmarshal([]byte(fileScanner.Text()), &readedList)
	}
	if err != nil {
		return fmt.Errorf("чтение данных из файла:: %w", err)
	}
	for i := range readedList.Tasks {
		list.AddTask(readedList.Tasks[i].Description, filename)
	}
	return nil
}

func WriteDataToFile(filename string, toDoList *ToDoList) error {
	f, err := os.OpenFile(filePath+filename, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("открытия файла: %w", err)
	}
	toDoListStr, err := json.Marshal(&toDoList)
	if err != nil {
		return fmt.Errorf("запись данных в файл: %w", err)
	}
	f.WriteString(string(toDoListStr))
	defer f.Close()
	return nil
}
