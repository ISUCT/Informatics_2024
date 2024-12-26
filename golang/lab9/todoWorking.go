package lab9

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Input() string {
	Reader := bufio.NewReader(os.Stdin)

	input, err := Reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	return input
}

func (l *ToDoList) AddExercise(description string, path string) {
	task := Exercise{description: description, status: "не выполнена"}
	for _, el := range l.Tasks {
		if el == task {
			fmt.Println("эта задача уже существует", description)
		}
	}
	l.Tasks = append(l.Tasks, task)
	WriteFile(path, l)
}

func (l *ToDoList) ShowTodoList() {
	if len(l.Tasks) == 0 {
		fmt.Println("Задач нет")
	}

	for i, task := range l.Tasks {
		fmt.Printf("%d. %s, %s\n", i+1, task.description, task.status)
	}
}

func (l *ToDoList) StatusComplete(number int) {
	if number < 1 || number > len(l.Tasks) {
		fmt.Printf("Задачи с таким номером нет")
		return
	}

	l.Tasks[number-1].status = "выполнена"
}

func (l *ToDoList) DeleteExercise(number int) {
	if number < 1 || number > len(l.Tasks) {
		fmt.Printf("Задачи с таким номером нет")
		return
	}

	l.Tasks = append(l.Tasks[:number-1], l.Tasks[number:]...)
}

func (l *ToDoList) SearchExercise(word string) string {
	for i := range l.Tasks {
		if strings.Contains(l.Tasks[i].description, word) {
			task, err := json.Marshal(l.Tasks[i])
			if err != nil {
				fmt.Println("Ошибка при кодировании файла: ", err)
			}
			return string(task)
		}
	}
	return ""
}
