package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"isuct.ru/informatics2022/lab4"
)

func recordingData(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		newLine, _ := reader.ReadString('\n')
		newLine = strings.TrimSpace(newLine)
		if newLine == "Нет" {
			fmt.Println("Данные успешно записаны")
			break
		}

		err := WriteFile(fileName, newLine)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func RunLab8() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите имя файла:")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	err := CreatingFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите текст который хотите записать в файл (Для завершения напишите 'Нет')")
	recordingData(fileName)

	file, err := ReadFile(fileName)
	if err != nil {
		panic("Произошла ошибка при чтении данных из файла")
	}
	fmt.Println("Содержимое файла: \n", file)

	fmt.Println("Введите текст для поиска в файле:")
	searchText, _ := reader.ReadString('\n')
	searchText = strings.TrimSpace(searchText)
	err = SearchingTextInFile(fileName, searchText)
	if err != nil {
		fmt.Println(err)
		return
	}

	variables, err := ReadFileForLab4(fileName)
	if err != nil {
		panic("Произошла ошибка при чтении данных из файла")
	}

	fmt.Print("Задача А\n")
	fmt.Print(lab4.TaskA(variables[1], variables[2], variables[3], variables[0]), "\n")
	fmt.Print("Задача В\n")
	fmt.Print(lab4.TaskB(variables[4:], variables[0]), "\n")
}
