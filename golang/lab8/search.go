package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchingTextInFile(fileName string, searchText string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Не удается открыть файл")
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	lineNumber := 1
	variable := false
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), searchText) {
			fmt.Printf("Текст найден на строке %d \n", lineNumber)
			variable = true
		}
		lineNumber++
	}
	if !variable {
		fmt.Println("Текст не найден")
	}
}
