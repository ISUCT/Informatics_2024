package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunLab8() {
	filename := "sample_file.txt"

	file := CreateFile(filename)
	WriteToFile(file)

	fmt.Println("Содержимое файла:")
	ReadFile(filename)

	fmt.Print("\nВведите текст для поиска: ")
	var searchText string
	fmt.Scanln(&searchText)
	SearchInFile(filename, searchText)
}

func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		os.Exit(1)
	}
	return file
}

func WriteToFile(file *os.File) {
	defer file.Close()
	fmt.Println("Введите текст для записи (завершите ввод пустой строкой):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		file.WriteString(line + "\n")
	}
	fmt.Println("Данные записаны в файл.")
}

func ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func SearchInFile(filename, searchText string) {
	file, _ := os.Open(filename)
	defer file.Close()

	found := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchText) {
			fmt.Println("Найдено:", line)
			found = true
		}
	}

	if !found {
		fmt.Println("Текст не найден.")
	}
}
