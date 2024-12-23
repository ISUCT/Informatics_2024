package lab8

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var errorFileAlreadyExists = errors.New("файл с таким именем уже существует")
var errSearchInFile = errors.New("файл не содержит данного текста")

func CreateFile(path string) error {
	_, errStat := os.Stat(path)
	if errStat == nil {
		return errorFileAlreadyExists
	}

	file, errCreate := os.Create(path)
	if errCreate != nil {
		return fmt.Errorf("(CreateFile) создание файла %s: %w", path, errCreate)
	}
	defer file.Close()

	return nil
}

func WriteFile(path string) error {
	file, errOpenFile := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if errOpenFile != nil {
		return fmt.Errorf("(WriteFile) открытие файла %s: %w", path, errOpenFile)
	}
	defer file.Close()

	text, err := ConsoleInput("Текст, который будет введён в файл")
	if err != nil {
		return fmt.Errorf("(WriteFile) ввод текста: %w", err)
	}

	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("(WriteFile) ошибка записи в файл %s: %w", path, err)
	}

	return nil
}

func ReadFile(path string) (string, error) {
	file, errOpen := os.Open(path)
	if errOpen != nil {
		return "", fmt.Errorf("(ReadFile) ошибка открытия файла %s: %w", path, errOpen)
	}
	defer file.Close()

	var result string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("(ReadFile) ошибка чтения файла %s: %w", path, err)
		}
		result += string(data[:n])
	}
	return result, nil
}

func SearchInFile(path string, searchText string) (bool, error) {
	StringFile, err := ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("(SearchInFile) ошибка чтения файла %s: %w", path, err)
	}

	if strings.Contains(StringFile, searchText) {
		return true, nil
	}
	return false, errSearchInFile
}

func ConsoleInput(prompt string) (string, error) {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Printf("Введите %s: ", prompt)
	text, err := in.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("(ConsoleInput) ошибка ввода: %w", err)
	}

	text = strings.TrimSpace(text)
	return text, nil
}
