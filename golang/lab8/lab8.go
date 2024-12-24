package lab8

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/lab4"
)

func Lab8() {
	path := "lab8/lab8.txt"

	errCreateFile := CreateFile(path)
	if errCreateFile != nil {
		log.Printf("(Lab8) создание файла: %v", errCreateFile)
	}

	text, err := ConsoleInput("Текст, который будет введён в файл")
	if err != nil {
		log.Printf("(WriteFile) ввод текста: %v", err)
	}

	err = WriteFile(text, path)
	if err != nil {
		log.Fatalf("(Lab8) запись файла: %v", err)
	}

	fmt.Println(ReadFile(path))

	searchText, errInput := ConsoleInput("текст для поиска")
	if errInput != nil {
		log.Fatalf("(Lab8) ввод текста для поиска: %v", errInput)
	}
	fmt.Print(SearchInFile(path, searchText))

	result, errReadFileForLab4 := ReadFileForLab4()
	if errReadFileForLab4 != nil {
		log.Fatal(errReadFileForLab4)
	}
	fmt.Println(lab4.CompleteTaskA(result[0], result[1], result[2], result[3], result[4]))
	fmt.Println(lab4.CompleteTaskB(result[0], result[1], result[5:]))
}
