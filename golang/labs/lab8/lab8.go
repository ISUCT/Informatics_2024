package lab8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CreateFile(filename string) (string, error) {
	if _, err := os.Stat(filename); err == nil {
		return "", fmt.Errorf("Файл %s уже существует", filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("Ошибка создания файла: %w", err)
	}
	defer file.Close()
	return filename, nil
}

func WriteToFile(filename string, info string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла: %w", err)
	}
	defer file.Close()
	_, err = io.WriteString(file, info)
	if err != nil {
		return fmt.Errorf("Ошибка при записи в файл: %w", err)
	}
	return nil
}

func ReadFile(filepath string) (string, error) {
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения файла: %w", err)
	}
	return string(fileData), nil
}

func SearchInFile(filedata string, searchText string) bool {
	return strings.Contains(filedata, searchText)
}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	return text
}

func RunLab8() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя файла: ")
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	filePath, err := CreateFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text := GetInput()

	info := fmt.Sprintf("%s\n", text)
	err = WriteToFile(filePath, info)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Данные успешно записаны в файл", filename)

	fileData, err := ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Содержимое файла:", fileData)

	var searchText string
	fmt.Print("Текст для поиска: ")
	searchText, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	searchText = strings.TrimSpace(searchText)

	if SearchInFile(fileData, searchText) {
		fmt.Println("Текст найден в файле")
	} else {
		fmt.Println("Текст не найден в файле")
	}
}
