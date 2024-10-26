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

	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Print("Введите задачу: ")
	newName, _ = in.ReadString('\n')
	newName = strings.Replace(newName, "\n", "", -1)

	for {
		fmt.Print("Введите дедлайн: ")
		fmt.Scan(&inputTime)

		i, err := strconv.Atoi(inputTime)
		if err == nil {
			newTime = i
			break
		}
		fmt.Printf("--дедлайн \"%v\" введён некорректно\n", inputTime)
	}

	fmt.Print("Введите тег: ")
	fmt.Scan(&newTeg)

	return newName, newTime, newTeg
}
