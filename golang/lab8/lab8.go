package lab8

import (
	"fmt"
	"log"
)

func CompleteLaba8() {
	path := "lab8/lab8.txt"

	errCreateFile := CreateFile(path)
	if errCreateFile != nil {
		log.Println(errCreateFile)
	}

	err := WriteFile(path)
	if err != nil {
		log.Fatalf("(CompleteLaba8) запись файла: %v", err)
	}

	// fmt.Print(ReadFile(path))
	searchText, errInput := InputText("текст для поиска")
	if errInput != nil {
		log.Fatalf("(CompleteLaba8) ввод текста для поиска: %v", errInput)
	}
	fmt.Print(SearchInFile(path, searchText))

	// result, errReadFileForLab4 := ReadFileForLab4()
	// if errReadFileForLab4 != nil {
	// 	log.Fatal(errReadFileForLab4)
	// }
	// fmt.Println(laba4.CompleteTaskA(result[0], result[1], result[2], result[3], result[4]))
	// fmt.Println(laba4.CompleteTaskB(result[0], result[1], result[5:]))
}
