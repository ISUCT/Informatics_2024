package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateFile(filename string) (string, error) {
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка при создание файла: %w", err)
	}

	defer file.Close()
	fmt.Printf("Файл %s успешно создан.", filename)
	return filename, nil
}

func WriteDataToFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0600)

	if err != nil {
		return fmt.Errorf("открытие файла: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "q" {
			break
		}
		_, err := file.WriteString(text)
		file.WriteString("\n")

		if err != nil {
			return fmt.Errorf("ошибка при записи в файл: %w", err)
		}
	}

	return nil
}

func ReadDataFromFile(filePath string) ([]string, error) {
	var textAll []string
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("открытие файла: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		textAll = append(textAll, scanner.Text())
	}

	return textAll, nil
}

func IsStrInFile(filePath, searchStr string) (bool, int, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return false, -1, fmt.Errorf("ошибка при открытие файла: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		if scanner.Text() == searchStr {
			return true, lineNumber, nil
		}
		lineNumber++
	}

	return false, -1, fmt.Errorf("строка %s не найдена", searchStr)
}

func RunLab8WorkWithFile() {
	var filename, strForFind string

	// Создание файла
	fmt.Println("Введите имя создаваемоего файла")
	fmt.Scan(&filename)
	file, err := CreateFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Запись данных в файл
	fmt.Println("Введите данные в файл (для завершения ввода введите q):")
	err = WriteDataToFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Чтение данных из файла
	fmt.Println(file)
	lines, err := ReadDataFromFile(file)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Чтение файла '%s' завершено. Содержимое:\n", filename)
	for _, line := range lines {
		fmt.Println(line)
	}

	// Поиск строки в файле
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку, которую хотите найти в файле: ")
	strForFind, _ = reader.ReadString('\n')
	strForFind = strings.TrimRight(strForFind, "\r\n")

	found, lineNumber, err := IsStrInFile(file, strForFind)
	if found {
		fmt.Printf("Строка '%s' найдена на %d строке.\n", strForFind, lineNumber+1)
	} else {
		fmt.Println(err)
	}
}
