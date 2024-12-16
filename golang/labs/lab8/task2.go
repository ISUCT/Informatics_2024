package lab8

import (
	"fmt"
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
			create()
		} else if mainChoose == 2 {
			write()
		} else if mainChoose == 3 {
			readFile()
		} else if mainChoose == 4 {
			searchingTextIntoFile()
		} else if mainChoose == 5 {
			break
		}
		fmt.Print("\n")
	}
}