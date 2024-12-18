package lab8

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func SearchText(filename string) {
	file, err := os.Open(filename)
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

func Lab8_2() {
	var text string
	file_name := "hello.txt"
	fmt.Println("Введите текст:")
	fmt.Fscan(os.Stdin, &text)
	file, err := os.Create(file_name)

	if err != nil {
		fmt.Println("Не удалось создать фаид:", err)
		os.Exit(1)
	}
	file.WriteString(text)
	fContent, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Введеная вами информация в файл:", string(fContent))
	SearchText("hello.txt")
}
