package lab9

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Task struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

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

func RunLab9Tasks() {
	name := StartProgram()
	number := 1
	fmt.Println("Вы можете: \n 1) добавить задачу \n 2) просмотреть все задачи")
	fmt.Println(" 3) отметить задачу как выполненную \n 4) удалить задачу \n 5) найти задачу \n 6) выйти из программы")
	for {
		var choise int
		fmt.Print("Введите номер задачи, который хотите выполнить: ")
		fmt.Fscan(os.Stdin, &choise)

		switch choise {
		case 1:
			var count int
			var arr []Task
			fmt.Print("Введите сколько задач Вы хотите ввести: ")
			fmt.Fscan(os.Stdin, &count)
			for i := 0; i < count; i++ {
				var task Task
				task.Number = number
				number++
				fmt.Print("Введите название задачи: ")
				fmt.Fscan(os.Stdin, &task.Name)
				task.Status = false
				arr = append(arr, task)
			}
			WriteTaskInJson(name, arr)
		case 2:
			Task := ReadTask(name)
			for _, task := range Task {
				if task.Status {
					fmt.Println("Задача №", task.Number, ":", task.Name, "- выполнена")
				} else {
					fmt.Println("Задача №", task.Number, ":", task.Name, "- не выполнена")
				}
			}
		case 3:
			var number int
			fmt.Print("Введите номер выполненной задачи: ")
			fmt.Fscan(os.Stdin, &number)
			Task := ReadTask(name)
			for index, task := range Task {
				if task.Number == number {
					Task[index].Status = true
				}
			}
			os.Truncate(name, 0)
			WriteTaskInJson(name, Task)
		case 4:
			var number int
			fmt.Print("Введите номер задачи, которую хотите удалить: ")
			fmt.Fscan(os.Stdin, &number)
			Task := ReadTask(name)
			for index, task := range Task {
				if task.Number == number {
					Task = append(Task[:index], Task[index+1:]...)
				}
			}
			os.Truncate(name, 0)
			WriteTaskInJson(name, Task)
		case 5:
			Task := ReadTask(name)
			var search string
			fmt.Print("Введите задачу, которую хотите найти: ")
			fmt.Fscan(os.Stdin, &search)
			fmt.Println("Если не выводит успех поиска, то задача не найдена")
			for _, task := range Task {
				if search == task.Name {
					fmt.Println("Задача найдена")
					break
				}
			}
		case 6:
			return
		}
	}
}
