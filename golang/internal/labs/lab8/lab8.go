package lab8

import "fmt"

func RunLab8() error {
	fmt.Println("Введите название файла")
	var filename string
	fmt.Scan(&filename)
	RunLab8forLab4(filename)
	fmt.Println("Введите название файла")
	fmt.Scan(&filename)
	f, err := CreateFile(filename)
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
	searchResult, err := SearchInFile(f, search)
	if err != nil {
		fmt.Println(err)
	}
	if len(searchResult) == 0 {
		fmt.Println("Ничего не найдено")
	} else {
		for _, result := range searchResult {
			fmt.Println(result)
		}
	}
	return err
}
