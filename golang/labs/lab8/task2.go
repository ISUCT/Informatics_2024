package lab8

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func ansTask2 () {
	for {
		var mainChoose int = 0
		
		fmt.Println("Что вы желаете сделать?")
		fmt.Println("1 - Создать файл/пересоздать")
		fmt.Println("2 - Записать в файл")
		fmt.Println("3 - Вывести файл в консоль")
		fmt.Println("4 - Найти текст в файле")
		fmt.Println("5 - Выйти из программы")
		fmt.Scan(&mainChoose)

		if mainChoose == 1 {
			var str string

			fmt.Print("Ведите название файла: ")
			fmt.Scan(&str)
			num := createFile(str)
			if num == 1 {
				var openChoose string
				fmt.Print("Желаете записать в этот файл что-нибудь?\nДа/нет ")
				fmt.Scan(&openChoose)
				if openChoose == "Да" {
					file, _ := os.OpenFile(str, os.O_RDWR, 0666)
					defer file.Close()
					writeToFile(file)
				} else {
					fmt.Print("\n")
					continue
				}
			
			} 
		} else if mainChoose == 2 {
				var str string

				fmt.Print("Введите название файла, в который хотите записать текст: ")
				fmt.Scan(&str)
				file, err := os.OpenFile(str, os.O_RDWR, 0666)
				if err != nil {
					if errors.Is(err, fs.ErrNotExist) {
						var createChoose string
						fmt.Print("Файла не существует, желаете создать?\nДа/нет ")
						fmt.Scan(&createChoose)
						if createChoose == "Да" {
							createFile(str)
						} else {
							fmt.Print("\n")
							continue
						}
					} else{
						fmt.Print("Не удалось открыть файл, ошибка: ", err)
						fmt.Print("\n")
						continue
					}
				}
				defer file.Close()
				writeToFile(file)
		} else if mainChoose == 3 {
			var str string

			fmt.Print("Введите название файла, который надо открыть: ")
			fmt.Scan(&str)
			read(str)
		} else if mainChoose == 4 {
			var strFile string
			var strSearch string


			fmt.Print("Введите название файла, в котором будет поиск: ")
			fmt.Scanln(&strFile)
			file, err := os.OpenFile(strFile, os.O_RDONLY, 0444)
				if err != nil {
					if errors.Is(err, fs.ErrNotExist) {
						fmt.Print("Файла не существует, выбирите другой файл")
					} else{
						fmt.Print("Не удалось открыть файл, ошибка:", err)
						fmt.Print("\n")
						continue
					}
				}
			defer file.Close()

			fmt.Print("Введите искомый текст: ")
			fmt.Scanln(&strSearch)

			search(file, strSearch)
		} else if mainChoose == 5 {
			break
		}
		fmt.Print("\n")
	}
}