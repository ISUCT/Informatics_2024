package lab9

import (
	"fmt"
	"os"
)

type Task struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
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
			ShowTask(Task)
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
			SearchTask(search, Task)
		case 6:
			return
		}
	}
}
