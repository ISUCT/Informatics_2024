package lab8

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func writeToFile (file *os.File) () {
	var str string
	var strBuf string
	fmt.Print("Введите желаемый текст (чтобы прекратить запись введите \"QAWSED\" после пробела): ")
	for {
		fmt.Scan(&str)
		if str == "QAWSED" {
			break
		}else {
			strBuf += " " + str
		}
	}
	_, err := file.WriteString(fmt.Sprint(strBuf[1:], "\n"))
	if err != nil {
		fmt.Println("Не удалось записать, ошибка:  ", err)
		return
	}
	fmt.Print("Запись успешна\n")
	return 
}

func checkingErr (err error, str string) {
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			var createChoose string
			fmt.Print("Файла не существует, желаете создать?\nДа/нет ")
			fmt.Scan(&createChoose)
			if createChoose == "Да" {
				createFile(str)
			} else {
				fmt.Print("\n")
			}
		} else{
			fmt.Print("Не удалось открыть файл, ошибка: ", err)
			fmt.Print("\n")
		}
	}
}

func write() {
	var str string

	fmt.Print("Введите название файла, в который хотите записать текст: ")
	fmt.Scan(&str)
	file, err := os.OpenFile(str, os.O_RDWR, 0666)
	checkingErr(err, str)
	defer file.Close()
	writeToFile(file)
}