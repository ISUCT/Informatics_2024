package lab8

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadDataFromFile(filename string) ([]string, error) {
	var textSlice []string
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.New("ошибка открытия файла")
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		textSlice = append(textSlice, fileScanner.Text())
	}
	if len(textSlice) == 0 {
		return nil, fmt.Errorf("файл пуст")
	}
	return textSlice, nil
}
