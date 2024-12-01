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

	errOpenKeyboard := keyboard.Open()
	if errOpenKeyboard != nil {
		log.Fatal(errOpenKeyboard)
	}
	defer keyboard.Close()

	for {
		clear()
		list.OutputTodo()

		text, _ := os.ReadFile(linkHelpCommand)
		fmt.Println(string(text))

		char, key, errGetKey := keyboard.GetKey()
		if errGetKey != nil {
			log.Fatal(errGetKey)
		}

		if char == '+' {
			clear()
			list.AddTask(console.InputTask())
		}

		if char == '-' {
			fmt.Println("введите номер задачи которую вы хотите удалить:")
			numRemove, errNumRemove := strconv.Atoi(console.Write())
			if errNumRemove != nil {
				log.Fatal(errNumRemove)
			}
			errRemoveTask := list.RemoveTask(numRemove)
			if errRemoveTask != nil {
				log.Fatal(errRemoveTask)
			}
		}

		if key == keyboard.KeyCtrlE {
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
		}

		if key == keyboard.KeyCtrlR {
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

			list.MoveTask(from, to)
		}

		if key == keyboard.KeyCtrlS {
			clear()
			lupSort(&list)
		}

		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlQ || char == 'q' || char == 'Q' || char == 'й' || char == 'Й' {
			break
		}

		clear()
	}

	list.Save()
}

func lupSort(list *todo.Todo) {
	status := '+'
	for {
		text, _ := os.ReadFile(linkHelpSorted)
		fmt.Println(string(text))
		fmt.Printf("\n\n\tctrl + R | обратная сортировка [%v]\n", string(status))

		char, key, errGetKey := keyboard.GetKey()
		if errGetKey != nil {
			log.Fatal(errGetKey)
		}

		if char == '1' {
			list.SortTaskNameAlphabeticalOrder()
			break
		}
		if char == '2' {
			list.SortTaskDataCreate()
			break
		}
		if char == '3' {
			list.SortTaskDeadline()
			break
		}

		if key == keyboard.KeyCtrlR {
			if status == '+' {
				status = ' '
			} else {
				status = '+'
			}
		}

		if status == '+' {
			list.SortReverse()
		}

		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlQ || char == 'q' || char == 'Q' || char == 'й' || char == 'Й' {
			break
		}
		clear()
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
