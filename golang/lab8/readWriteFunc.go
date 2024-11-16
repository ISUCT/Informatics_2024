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

func CreateFile(path string) error {
	_, errStat := os.Stat(path)
	if errStat == nil {
		return errorFileAlreadyExists
	}

	file, errCreate := os.Create(path)
	if errCreate != nil {
		return errCreate
	}
	file.Close()

	return nil
}

func WriteFile(path string) error {
	file, errOpenFile := os.OpenFile(path, os.O_WRONLY, 0666)
	if errOpenFile != nil {
		return errOpenFile
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
