package lab8

import (
	"fmt"
	"os"
)

func CreatingFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Не удается открыть файл")
		panic(err)
	}
	defer file.Close()
}
