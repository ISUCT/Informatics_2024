package laba8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const Path = "lab8/text.txt"

func CreateFile() error {
	file, err := os.Create(Path)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}

func WriteFile() error {
	file, err := os.OpenFile(Path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("введите текст")
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	text, err := in.ReadString('\n')
	if err != nil {
		return err
	}

	file.WriteString(text)

	return nil
}

func ReadFile() (string, error) {
	file, err := os.Open(Path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var text string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		text = string(data[:n])
	}
	return text, nil
}

func SearchText(searchText string) (bool, error) {
	text, err := ReadFile()
	if err != nil {
		return false, err
	}
	search := strings.Contains(text, searchText)
	return search, nil
}
