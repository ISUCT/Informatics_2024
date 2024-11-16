package lab8

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/laba4"
)

func CompleteLaba8() {
	path := "lab8/lab8.txt"

	errCreateFile := CreateFile(path)
	if errCreateFile != nil {
		log.Println(errCreateFile)
	}

	errWriteFile := WriteFile(path)
	if errWriteFile != nil {
		log.Fatal(errWriteFile)
	}

	fmt.Print(ReadFile(path))

	result, errReadFileForLab4 := ReadFileForLab4()
	if errReadFileForLab4 != nil {
		log.Fatal(errReadFileForLab4)
	}
	fmt.Println(laba4.CompleteTaskA(result[0], result[1], result[2], result[3], result[4]))
	fmt.Println(laba4.CompleteTaskB(result[0], result[1], result[5:]))
}
