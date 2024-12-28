package lab8

import (
	"bufio"
	"fmt"
	"os"
)

const path = "labs/lab8/text.txt"

func RunLab8() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var err error
	err = CreateFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Введите текст который будет введён")
	TextForWrite, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	err = WriteFile(path, TextForWrite)
	if err != nil {
		panic(err)
	}

	text, err := ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

	fmt.Println("Введите текст для поиска")
	TextForSearch, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	textFound, err := SearchText(path, TextForSearch)
	if err != nil {
		panic(err)
	}
	fmt.Println(textFound)

	Task1()
}
