package lab8

import (
	"fmt"
	"os"
	"strings"
)

func CreateFile() (string, error) {
	filename := "text.txt"
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("Файл существует")
		return filename, nil
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("Ошибка существования файла: %w", err)
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer file.Close()
	return filename, nil
}
func WriteToFile(filepath string, name string, age int, city string, university string, additionalText string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла для записи: %w", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Имя: %s\nВозраст: %d\nГород: %s\nУниверситет: %s\nДополнительный текст: %s\n", name, age, city, university, additionalText)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}
	return nil
}

func ReadFile(filepath string) (string, error) {
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return string(fileData), nil
}
func SearcInFile(filedata string, searchText string) bool {
	return strings.Contains(filedata, searchText)
}

func RunFileLab() {
	filePath, err := CreateFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var name string
	var age int
	var city string
	var university string
	var additionalText string

	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	fmt.Print("Введите город: ")
	fmt.Fscan(os.Stdin, &city)
	fmt.Print("Введите институт, в котором вы учитесь: ")
	fmt.Fscan(os.Stdin, &university)
	fmt.Print("Введите любой текст: ")
	fmt.Scanln()
	fmt.Scanln(&additionalText)

	err = WriteToFile(filePath, name, age, city, university, additionalText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Данные успешно записаны в файл text.txt")

	fileData, err := ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Содержимое файла:", fileData)

	var searchText string
	fmt.Print("Введите текст для поиска: ")
	fmt.Scanln(&searchText)

	if SearcInFile(fileData, searchText) {
		fmt.Println("Текст найден в файле")
	} else {
		fmt.Println("Текст не найден в файле")
	}
}
