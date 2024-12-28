package lab8

import (
	"io"
	"os"
	"strings"
)

func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func WriteFile(path, text string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(text)
	return nil
}

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
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

func SearchText(path, searchText string) (bool, error) {
	text, err := ReadFile(path)
	if err != nil {
		return false, err
	}

	search := strings.Contains(text, searchText)
	return search, nil
}
