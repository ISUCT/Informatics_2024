package lab8

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var file_name string

func GetName() {
	fmt.Println("Введите имя файла:")
	fmt.Fscan(os.Stdin, &file_name)
	file_name += ".txt"
}

func SearchText() {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()
	fmt.Println("Введите текст для поиска в файле:")
	var searchText string
	fmt.Fscan(os.Stdin, &searchText)

	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchText) {
			fmt.Println("Текст был успешно найден:", line)
			found = true
		}
	}
	if !found {
		fmt.Println("Не удалось найти текст в файле")
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}
}

func CreateFile() {
	os.Create(file_name)
}

func WriteText() {
	var text string
	fmt.Println("Введите текст:")
	fmt.Fscan(os.Stdin, &text)
	file, err := os.OpenFile(file_name, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()
	if _, err := file.WriteString(text); err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func ReadText() {
	fContent, err := ioutil.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Введеная вами информация в файл:", string(fContent))
}
