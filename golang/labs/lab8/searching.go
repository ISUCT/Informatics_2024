package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchTextInFile(fileName, searchText string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("не удается открыть файл: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	lineNumber := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchText) {
			fmt.Printf("Текст найден на строке %d: %s\n", lineNumber, scanner.Text())
			found = true
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка сканирования файла: %w", err)
	}

	if !found {
		fmt.Println("Текст не найден.")
	}

	return nil
}
