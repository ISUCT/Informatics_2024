package lab9

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func StartProgram() string {
	var name string
	fmt.Print("Введите название файла с Вашими задачами с окончанием .json: ")
	fmt.Fscan(os.Stdin, &name)
	file, _ := os.OpenFile(name, os.O_RDWR, 0666)
	defer file.Close()
	return name
}

func WriteTaskInJson(name string, arr []Task) {
	file, _ := os.OpenFile(name, os.O_RDWR, 0666)
	defer file.Close()
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(arr); err != nil {
		return
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(buf.String())
	writer.Flush()
}

func ReadTask(name string) []Task {
	file, _ := os.Open(name)
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	var Task []Task
	json.Unmarshal(data, &Task)
	return Task
}

func ShowTask(tasks []Task) {
	for _, task := range tasks {
		if task.Status {
			fmt.Println("Задача №", task.Number, ":", task.Name, "- выполнена")
		} else {
			fmt.Println("Задача №", task.Number, ":", task.Name, "- не выполнена")
		}
	}
}

func SearchTask(search string, tasks []Task) {
	for _, task := range tasks {
		if search == task.Name {
			fmt.Println("Задача найдена")
			break
		}
	}
}
