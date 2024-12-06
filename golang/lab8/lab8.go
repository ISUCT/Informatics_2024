package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunLab8() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите имя файла:")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	CreatingFile(fileName)

	fmt.Println("Введите текст который хотите записать в файл (Для завершения напишите 'Нет')")
	for {
		newLine, _ := reader.ReadString('\n')
		newLine = strings.TrimSpace(newLine)
		if newLine == "Нет" {
			fmt.Println("Данные успешно записаны")
			break
		}
		WriteFile(fileName, newLine)
	}

	variables := ReadFileForLab4(fileName)
	fmt.Printf("Содержание файла: %.2f \n", variables[0:])

	fmt.Println("Введите текст для поиска в файле:")
	searchText, _ := reader.ReadString('\n')
	searchText = strings.TrimSpace(searchText)
	SearchingTextInFile(fileName, searchText)

	fmt.Print("Задача А\n")
	fmt.Print(TaskA(variables[0], variables[1], variables[2], variables[3]), "\n")
	fmt.Print("Задача В\n")
	fmt.Print(TaskB(variables[4:], variables[0]), "\n")
}
