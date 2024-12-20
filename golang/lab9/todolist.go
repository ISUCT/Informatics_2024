package lab9

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func (t *ToDoList) AddTask(fileName string, description string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("открытия файла: %w", err)
	}
	defer file.Close()

	task := ToDoTask{Description: description, Status: "не выполнена"}
	t.Tasks = append(t.Tasks, task)
	newLine, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("ошибка при переводе файла JSON: %w", err)
	}

	file.WriteString(string(newLine))
	return nil
}

func (t *ToDoList) PrintToDoList() {
	if len(t.Tasks) == 0 {
		fmt.Println("Нет добавленных задач")

	}
	for i, task := range t.Tasks {
		fmt.Printf("Задачa %d. Описание задачи: %s, Статус задачи: %s\n", i+1, task.Description, task.Status)
	}
}

func (t *ToDoList) StatusСhange(index int) {
	if index < 1 || index > len(t.Tasks) {
		fmt.Print("Нет задачи с таким номером")
		return
	}

	t.Tasks[index-1].Status = "выполнена"
	fmt.Printf("Статус для %v задачи обнавлён\n", index)
}

func (t *ToDoList) DeletingTask(index int) {
	if index < 1 || index > len(t.Tasks) {
		fmt.Print("Нет задачи с таким номером")
		return
	}

	t.Tasks = append(t.Tasks[:index-1], t.Tasks[index:]...)
	fmt.Printf("Задача %v удалена\n", index)
}

func (t *ToDoList) TaskSearch(keyWord string) (string, error) {
	for i := range t.Tasks {
		if strings.Contains(t.Tasks[i].Description, keyWord) {
			task, err := json.Marshal(t.Tasks[i])
			if err != nil {
				return "", fmt.Errorf("ошибка при переводе файла JSON: %w", err)
			}
			return string(task), nil
		}
	}
	return "", nil
}
