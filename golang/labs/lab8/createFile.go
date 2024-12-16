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