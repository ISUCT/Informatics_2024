package lab8

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var errorFileAlreadyExists = errors.New("фаил с таким именем уже существует")
var errSearchInFile = errors.New("файл не содержит искомого текста")

func CreateFile(path string) error {
	_, errStat := os.Stat(path)
	if errStat == nil {
		return errorFileAlreadyExists
	}

	file, errCreate := os.Create(path)
	if errCreate != nil {
		return fmt.Errorf("(CreateFile) создание файла %s: %w", path, errCreate)
	}
	file.Close()

	return nil
}

func WriteFile(path string) error {
	file, errOpenFile := os.OpenFile(path, os.O_WRONLY, 0666)
	if errOpenFile != nil {
		return fmt.Errorf("(WriteFile) открытие файла %s: %w", path, errOpenFile)
	}
	defer file.Close()

	text, err := InputText("текст который будет введён в файл")
	if err != nil {
		return fmt.Errorf("(WriteFile) ввод: %w", err)
	}

	file.WriteString(text)

	return nil
}

func ReadFile(path string) (string, error) {
	file, errOpen := os.Open(path)
	if errOpen != nil {
		return "", errOpen
	}
	defer file.Close()

	var result string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		result = string(data[:n])
	}
	return result, nil
}

func SearchInFile(path string, searchText string) (bool, error) {
	StringFile, err := ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("(SearchInFile)ReadFile %s: %w", path, err)
	}

	return strings.Contains(string(StringFile), searchText), nil
}

func InputText(text string) (string, error) {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Printf("Введите %v: ", text)
	text, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	return text, nil
}
