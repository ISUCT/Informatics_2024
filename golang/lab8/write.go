package lab8

import (
	"fmt"
	"os"
)

func WriteFile(fileName string, newLine string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Не удается открыть файл")
		panic(err)
	}
	defer file.Close()

	file.WriteString(newLine + "\n")
}
