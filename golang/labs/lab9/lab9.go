package lab9

import (
	"bufio"
	"fmt"
	"os"
)

const filePath = "../golang/labs/lab9/"

var command string

var filename string

func RunLab9() {
	fmt.Println("Введите имя файла, в который вы хотите записать данные:")
	fmt.Scan(&filename)

	name, err := CreateFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	TaskList := ToDoList{}

	err = ReadDataFromFile(name, &TaskList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Для дальнейшего действия введите одну из приведенных ниже команд:\n" +
		"1. Добавить задачу\n" +
		"2. Показать все задачи\n" +
		"3. Отметить задачу как выполненную\n" +
		"4. Удалить задачу\n" +
		"5. Поиск задачи\n" +
		"6. Выйти")
	for {
		fmt.Println("Ожидание ввода команды:")
		fmt.Scan(&command)
		switch command {
		case "1":
			fmt.Println("Введите описание задачи:")
			_, _ = reader.ReadString('\n')
			newTask, _ := reader.ReadString('\n')
			newTask = newTask[:len(newTask)-2]
			err = TaskList.AddTask(newTask, name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Задача '%s' добавлена\n", newTask)
		case "2":
			TaskList.PrintToDoList()
		case "3":
			fmt.Println("Введите ключевое слово для поиска задачи, которую хотите пометить как выполненную:")
			var description string
			fmt.Scan(&description)
			taskName, err := TaskList.MarkAsComplete(description, name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Задача '%s' выполнена\n", taskName)
		case "4":
			fmt.Println("Введите индекс задачи, которую вы хотите удалить:")
			var index int
			fmt.Scan(&index)
			fmt.Println(index)
			taskName, err := TaskList.DeleteTask(index, name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Задача '%s' удалена\n", taskName)
		case "5":
			fmt.Println("Введите ключевое слово для поиска задачи:")
			var taskName string
			fmt.Scan(&taskName)
			searchedTask, err := TaskList.SearchTask(taskName)
			if err == nil {
				fmt.Println(searchedTask)
			}
		case "6":
			fmt.Println("Работа программы завершена!")
			return
		default:
			fmt.Println("Неверно введена команда")
		}
	}
}
