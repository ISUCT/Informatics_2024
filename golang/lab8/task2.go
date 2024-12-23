package lab8

import (
	"fmt"
	"os"
)

func FileCreation() {
	file, err := os.Create("Task2.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл")
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())
	written, err := file.WriteString("Запись текста в файл")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Записано %#v байт в новый файл\n", written)
}
