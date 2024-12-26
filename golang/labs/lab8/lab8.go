package lab8

import (
	"fmt"

	fileutilis "isuct.ru/informatics2022/labs/lab8/fileutilis"
)

func RunLab8() {
	var filename string
	fmt.Print("Введите название файла: ")
	fmt.Scan(&filename)
	var info string
	fmt.Print("Введите информацию для записи в файл: ")
	fmt.Scan(&info)
	var searchString string
	fmt.Print("Введите слово для поиска: ")
	fmt.Scan(&searchString)
	_, err := fileutilis.CreateFile(filename)
	if err != nil {
		fmt.Printf("Ошибка создания файла: %v\n", err)
		return
	}
	err = fileutilis.WriteToFile(filename, info)
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
		return
	}
	result, err := fileutilis.SearchWord(filename, searchString)
	if err != nil {
		fmt.Printf("Ошибка поиска слова: %v\n", err)
		return
	}
	fmt.Println(result)
	fileContent, err := fileutilis.ReadFromFile(filename)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}
	fmt.Println("Содержимое файла:\n", fileContent)
}
