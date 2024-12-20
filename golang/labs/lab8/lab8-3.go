package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
