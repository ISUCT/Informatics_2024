package lab8

import (
	"bufio"
	"fmt"
	"os"
)

func ReadDataFromFile(filename string) ([]string, error) {
	var textSlice []string
	f, err := os.OpenFile(filePath+filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("открытие файла: %w", err)
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		textSlice = append(textSlice, fileScanner.Text())
	}
	if len(textSlice) == 0 {
		return nil, fmt.Errorf("чтение файла: файл пуст")
	}
	return textSlice, nil
}
