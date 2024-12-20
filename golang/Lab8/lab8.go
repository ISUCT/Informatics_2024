package lab8

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func RunLab8() {
	err := CreateFile()
	if err != nil {
		panic(err)
	}

	err = WriteFile()
	if err != nil {
		panic(err)
	}
	text, err := ReadFile(Linc)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

	err = SearchFile()
	if err != nil {
		panic(err)
	}

	input, err := ReadFile(LincTask1)
	if err != nil {
		panic(err)
	}
	list, err := ChangeStringToNumber(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(lab4.TaskA(list[0], list[1], list[2], list[3], list[4]))

	fmt.Println(lab4.TaskB(list[0], list[1], list[5:]))
}
