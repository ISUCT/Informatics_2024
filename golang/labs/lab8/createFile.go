package lab8

import (
	"fmt"
	"os"
)

func createFile(str string) (int) {
	file, err := os.Create(str)
	if err != nil {
		fmt.Println("Не удалось создать файл, ошибка: ", err)
		return 0
	}
	fmt.Print("Файл успешно создан\n")
	defer file.Close()
	return 1
}

func create() {
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
		}
	} 
}