package lab8

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var errorFileAlreadyExists = errors.New("фаил с таким именем уже существует")
var errSearchInFile = errors.New("файл не содержит искомого текста")

func CreateFile(path string) error {
	_, errStat := os.Stat(path)
	if errStat == nil {
		return errorFileAlreadyExists
	}

	file, errCreate := os.Create(path)
	if errCreate != nil {
		return fmt.Errorf("(CreateFile) создание файла %s: %w", path, errCreate)
	}
	file.Close()

	return nil
}

func WriteFile(path string) error {
	file, errOpenFile := os.OpenFile(path, os.O_WRONLY, 0666)
	if errOpenFile != nil {
		return fmt.Errorf("(WriteFile) открытие файла %s: %w", path, errOpenFile)
	}
	defer file.Close()

	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Print("Введите текст: ")
	text, _ := in.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	file.WriteString(text)

	return nil
}

func ReadFile(path string) string {
	file, errOpen := os.Open(path)
	if errOpen != nil {
		log.Fatal(errOpen)
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
	return result
}

func SearchInFile(path string, searchText string) (int, error) {
	var n int = 0
	var i int = 0
	file := ReadFile(path)

	for _, f := range file {
		if f == '\n' {
			n++
		}
		if byte(f) == searchText[i] {
			i++
			return n, nil
		} else {
			i = 0
		}
	}
	return 0, errSearchInFile
}
