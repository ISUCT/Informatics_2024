package lab8

import (
	"bufio"
	"fmt"
	"os"
)

func RunLab8() {
	var err error

	err = CreateFile()
	if err != nil {
		panic(err)
	}

	fmt.Println("Введите текст")
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	text, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	err = WriteFile(Path, text)
	if err != nil {
		panic(err)
	}

	text, err = ReadFile(Path)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

	err = task1()
	if err != nil {
		panic(err)
	}
}
