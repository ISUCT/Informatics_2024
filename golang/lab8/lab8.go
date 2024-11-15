package lab8

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CompleteLaba8() {
	CreateFile("lab8.txt")
	WriteFile()
	ReadFile()
}

func CreateFile(name string) {
	_, errStat := os.Stat(name)
	if errStat != nil {
		os.Exit(1)
		log.Fatal(errStat)
	}

	file, errCreate := os.Create(name)
	if errCreate != nil {
		os.Exit(1)
		log.Fatal(errCreate)
	}
	file.Close()
}

func WriteFile() {
	file, err := os.Open("hello.txt")
	if err != nil {
		os.Exit(1)
		log.Fatal(err)
	}
	defer file.Close()

	//ввод текста
	var text string = "текст"

	file.WriteString(text + " ")
}

func ReadFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		os.Exit(1)
		log.Fatal(err)
	}
	defer file.Close()

	var result string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		result = string(data[:n])
	}
	fmt.Print(result)
}
