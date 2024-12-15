package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunLab8() {
	filename := "output.txt"

	CreateAndWriteFile(filename)

	fmt.Println("\nСодержимое файла:")
	ReadFile(filename)

	fmt.Print("\nВведите текст для поиска: ")
	var searchText string
	fmt.Scanln(&searchText)
	SearchInFile(filename, searchText)
}

func CreateAndWriteFile(filename string) {
	file, _ := os.Create(filename)
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
	file, _ := os.Open(filename)
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
