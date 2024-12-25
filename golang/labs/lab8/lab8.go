package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunLab8() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя файла: ")
	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка ввода: %v\n", err)
		return
	}
	fileName = strings.TrimSpace(fileName)

	if err := CreateFile(fileName); err != nil {
		fmt.Println(err)
		return
	}

	var inputLines []string
	for {
		fmt.Println("Введите строку для записи (или 'Нет' для завершения):")
		newLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Ошибка чтения ввода: %v\n", err)
			return
		}
		newLine = strings.TrimSpace(newLine)
		if strings.EqualFold(newLine, "Нет") {
			break
		}
		inputLines = append(inputLines, newLine)
	}

	if err := AppendDataToFile(fileName, inputLines); err != nil {
		fmt.Println(err)
		return
	}

	content, err := ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Содержимое файла:\n%s\n", content)

	fmt.Print("Введите текст для поиска в файле: ")
	searchText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка ввода: %v\n", err)
		return
	}
	searchText = strings.TrimSpace(searchText)

	if err := SearchTextInFile(fileName, searchText); err != nil {
		fmt.Println(err)
	}
}
