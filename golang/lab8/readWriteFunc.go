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

var errorFileAlreadyExists = errors.New("фаил с таким же именем уже есть")
var errSearchInFile = errors.New("в файле нету нужного текста")

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

	text, err := InputText("текст который будет в файле")
	if err != nil {
		return fmt.Errorf("(WriteFile) ввод: %w", err)
	}

	file.WriteString(text)

	return nil
}

func ReadFile(path string) (string, error) {
	file, errOpen := os.Open(path)
	if errOpen != nil {
		return "", errOpen
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
	return result, nil
}

func SearchInFile(path string, searchText string) (int, error) {
	var lineNum int = 1
	var sign int = 0
	StringFile, err := ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("(SearchInFile)ReadFile %s: %w", path, err)
	}

	for _, f := range StringFile {
		log.Printf("N символ: %v, N строка: %v\n чтение: %v, поиск: %v\n", sign, lineNum, string(f), string(searchText[sign]))
		if byte(f) == '\n' {
			lineNum++
		}
		if byte(f) == searchText[sign] {
			log.Println("\t", f, "=", searchText[sign])
			sign++
		} else {
			log.Println("\t", f, "!=", searchText[sign])
			sign = 0
		}
		if sign == len(searchText) {
			return lineNum, nil
		}
	}
	return 0, errSearchInFile
}

func InputText(text string) (string, error) {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Printf("Введите %v: ", text)
	text, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	return text, nil
}
//Ответ на комментарий по строкам 38-41("По названию функции я ожидаю, что она будет записывать что-то в файл.
//Т.е. я ожидаю, что на входе будет две переменных: название файла и данные. Пожалуйста, обновите эту функцию чтобы она соответствовала неймингу"):

//Может я что не понимаю, но вроде бы ввожу текст с консоли, который будет записан в файл вот этой функцией:
//func InputText(text string) (string, error) {
//var in *bufio.Reader = bufio.NewReader(os.Stdin)

//fmt.Printf("Введите %v: ", text)
//text, err := in.ReadString('\n')
//if err != nil {
//return "", err
//}
//text = strings.Replace(text, "\n", "", -1)
//text = strings.Replace(text, "\r", "", -1)
//return text, nil
//}
