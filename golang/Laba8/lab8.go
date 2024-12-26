package lab8

import (
	"fmt"
)

func Laba8() {
	var err error

	err = CreateFile()
	if err != nil {
		panic(err)
	}
	err = WriteFile()
	if err != nil {
		panic(err)
	}
	text, err := ReadFile()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

	err = task1()
	if err != nil {
		panic(err)
	}
}
