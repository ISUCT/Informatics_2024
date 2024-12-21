package lab9

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/eiannone/keyboard"
	"isuct.ru/informatics2022/todolist/console"
	"isuct.ru/informatics2022/todolist/todo"
)

const (
	linkHelpSorted  = "todolist/help_sorted"
	linkHelpCommand = "todolist/help_command"
)

func CompleteLab9() {
	var list todo.Todo

	errLoad := list.Load()
	if errLoad != nil {
		log.Fatal(errLoad)
	}
	defer list.Save()

	errOpenKeyboard := keyboard.Open()
	if errOpenKeyboard != nil {
		log.Fatal(errOpenKeyboard)
	}
	defer keyboard.Close()

	for {
		clear()
		list.OutputTodo()
		close := interactionPanel(&list)
		if close {
			break
		}
	}
}

func interactionPanel(list *todo.Todo) bool {
	text, _ := os.ReadFile(linkHelpCommand)
	fmt.Println(string(text))

	char, _, errGetKey := keyboard.GetKey()
	if errGetKey != nil {
		log.Fatal(errGetKey)
	}

	switch char {
	case '+', '=':
		clear()
		list.AddTask(console.InputTask())
	case '-':
		fmt.Println("введите номер задачи которую вы хотите удалить:")
		numRemove, errNumRemove := strconv.Atoi(console.Write())
		if errNumRemove != nil {
			log.Fatal(errNumRemove)
		}

		errRemoveTask := list.RemoveTask(numRemove)
		if errRemoveTask != nil {
			log.Fatal(errRemoveTask)
		}
	case 'E', 'e', 'У', 'у':
		fmt.Println("введите номер задачи которую нужно изменить:")
		Change, errChange := strconv.Atoi(console.Write())
		if errChange != nil {
			log.Fatal(errChange)
		}

		clear()
		Name, Time, Teg := console.InputTask()
		errChangeTask := list.EditTask(Change, Name, Time, Teg)
		if errChangeTask != nil {
			log.Fatal(errChangeTask)
		}
	case 'R', 'r', 'К', 'к':
		interactionMove(list)
	case 'S', 's', 'В', 'в':
		status := ' '
		for {
			close := interactionSort(list, &status)
			if close {
				break
			}
		}
		if status == '+' {
			list.SortReverse()
		}
	case 'T', 't', 'Е', 'е':
		clear()
		fmt.Print("Введите текст для поиска: ")
		list.SerthTask(console.Write())
		interactionPanel(list)
	case 'Q', 'q', 'Й', 'й':
		return true
	default:
		return false
	}

	return false
}

func interactionMove(list *todo.Todo) {
	clear()
	list.OutputTodo()

	fmt.Println("введите откуда-куда переместить задачу")
	fmt.Print("откуда:")
	from, errFrom := strconv.Atoi(console.Write())
	if errFrom != nil {
		log.Fatal(errFrom)
	}

	fmt.Print("куда:")
	to, errTo := strconv.Atoi(console.Write())
	if errTo != nil {
		log.Fatal(errTo)
	}

	errMoveTask := list.MoveTask(from, to)
	if errMoveTask != nil {
		log.Fatalf("(interactionPanel) %v", errMoveTask)
	}
}

func interactionSort(list *todo.Todo, status *rune) bool {
	clear()
	text, _ := os.ReadFile(linkHelpSorted)
	fmt.Println(string(text))
	fmt.Printf("\n\n\tctrl + R | обратная сортировка [%v]\n", string(*status))

	char, _, errGetKey := keyboard.GetKey()
	if errGetKey != nil {
		log.Fatal(errGetKey)
	}

	switch char {
	case '1':
		list.SortTaskNameAlphabeticalOrder()
	case '2':
		list.SortTaskDataCreate()
	case '3':
		list.SortTaskDeadline()
	case 'R', 'r', 'К', 'к':
		if *status == '+' {
			*status = ' '
		} else {
			*status = '+'
		}
	case 'Q', 'q', 'Й', 'й':
		return true
	default:
		return false
	}

	return true
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
