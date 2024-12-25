package lab8

import (
	"fmt"
	"log"
)

func CompleteLab8() {
	path := "lab8/lab8.txt"

	errCreateFile := CreateFile(path)
	if errCreateFile != nil {
		log.Println(errCreateFile)
	}

	err := WritelnFile(path)
	if err != nil {
		log.Fatalf("(Completelab8) запись файла: %v", err)
	}

	searchText, errEnter := EnterText("текст для поиска")
	if errEnter != nil {
		log.Fatalf("(Completelab8) ввод текста для поиска: %v", errInput)
	}
	fmt.Print(SearchInFile(path, searchText))
}
