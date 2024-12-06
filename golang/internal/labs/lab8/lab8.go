package lab8

import "fmt"

func RunLab8() error {
	filename := "./internal/labs/lab8/data.txt"
	RunLab8forLab4(filename)
	filename = "./internal/labs/lab8/input.txt"
	_, err := CreateFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = WriteToFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = ReadFromFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	var search string
	fmt.Println("Введите текст для поиска: ")
	fmt.Scan(&search)
	err = SearchInFile(filename, search)
	return err
}
