package lab8

import (
	"fmt"
)

func CompleteLab8() {
	var err error

	err = CreateFile("lab8/text.txt")
	if err != nil {
		panic(err)
	}

	err = WriteFile("lab8/text.txt")
	if err != nil {
		panic(err)
	}

	text, err := ReadFile("lab8/text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

	err = task1()
	if err != nil {
		panic(err)
	}
}
