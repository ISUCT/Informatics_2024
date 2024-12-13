package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchingTextInFile(fileName string, searchText string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("не удается открыть файл")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	lineNumber := 1
	found := false
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), searchText) {
			fmt.Printf("Текст найден на строке %d \n", lineNumber)
			found = true
		}
		lineNumber++
	}

	if !found {
		fmt.Println("Текст не найден")
	}

	return err
}
