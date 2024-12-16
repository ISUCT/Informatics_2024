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
	found := false
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), searchText) {
			fmt.Print("Текст найден на строке \n")
			found = true
		}
	}

	if !found {
		fmt.Println("Текст не найден")
	}

	return err
}
