package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputTask() (string, int, string) {
	var newName string = "empty"
	var newTime int = 0
	var newTeg string = "empty"

	var inputTime string

	fmt.Print("Введите задачу: ")
	newName = Write()

	fmt.Print("Введите дедлайн: ")
	inputTime = Write()

	i, err := strconv.Atoi(inputTime)
	if err != nil {
		fmt.Printf("--дедлайн \"%v\" введён некорректно\n", inputTime)
	}
	newTime = i

	fmt.Print("Введите тег: ")
	newTeg = Write()

	return newName, newTime, newTeg
}

func Write() string {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}
