package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Linc      = "lab8/text.txt"
	WriteText = "Введите текст"
)

func CreateFile() error {
	file, err := os.Create(Linc)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func WriteFile() error {
	file, err := os.OpenFile(Linc, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Println(WriteText)
	text, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	file.WriteString(text)
	return nil
}

func ReadFile(linc string) (string, error) {
	file, err := os.Open(linc)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data := make([]byte, 64)

	n, err := file.Read(data)

	if err != nil {
		return "", err
	}

	return string(data[:n]), nil

}

func SearchFile() error {
	text, err := ReadFile(Linc)
	if err != nil {
		return err
	}
	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Println(WriteText)
	searchText, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	status := strings.Contains(text, searchText)
	fmt.Println(status)
	return nil
}
