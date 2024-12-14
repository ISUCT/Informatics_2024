package lab8

import (
	"fmt"

	fileutilis "isuct.ru/informatics2022/labs/lab8/fileutilis"
)

func RunLab8() {
	filename, err := fileutilis.CreateFile()
	if err != nil {
		fmt.Printf("Ошибка создания файла: %v\n", err)
		return
	}

	err = fileutilis.WriteToFile(filename)
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
		return
	}

	result, err := fileutilis.SearchWord(filename)
	if err != nil {
		fmt.Printf("Ошибка при поиске слова: %v\n", err)
		return
	}
	fmt.Println(result)

	fileContent, err := fileutilis.ReadFromFile(filename)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		return
	}
	fmt.Println("Содержимое файла:\n", fileContent)
}
