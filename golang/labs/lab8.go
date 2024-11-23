package labs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println("Файл успешно создан:", fileName)
	return nil
}

func writeFile(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		return err
	}
	fmt.Println("Данные успешно записаны в файл.")
	return nil
}

func readFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Содержимое файла:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}

func searchInFile(fileName string, searchText string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Результаты поиска:")
	scanner := bufio.NewScanner(file)
	found := false
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchText) {
			fmt.Printf("Найдено на строке %d: %s\n", lineNumber, line)
			found = true
		}
		lineNumber++
	}

	if !found {
		fmt.Println("Текст не найден.")
	}

	return scanner.Err()
}

func RunLab8() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите имя файла:")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	// Использование fileName для создания файла
	err := createFile(fileName)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}

	// Ввод текста для записи в файл
	fmt.Println("Введите текст для записи в файл (enter для завершения):")
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			break
		}
		err = writeFile(fileName, input)
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
	}

	// Чтение из файла
	err = readFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Поиск текста в файле
	fmt.Println("Введите текст для поиска в файле:")
	searchText, _ := reader.ReadString('\n')
	searchText = strings.TrimSpace(searchText)
	err = searchInFile(fileName, searchText)
	if err != nil {
		fmt.Println("Ошибка поиска в файле:", err)
	}
}

