package lab8

import (
	"fmt"
	"os"
)

func writeToFile (file *os.File) (int) {
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
		return 0
	}
	fmt.Print("Запись успешна\n")
	return 1
}