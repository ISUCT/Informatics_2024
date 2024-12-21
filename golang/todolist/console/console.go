package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputTask() (string, string, string) {
	fmt.Print("Введите задачу: ")
	newName := Write()

	fmt.Print("Введите дедлайн в формате (2006-01-02): ")
	newTime := Write()

	fmt.Print("Введите тег: ")
	newTeg := Write()

	return newName, newTime, newTeg
}

func Write() string {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		return ""
	}
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}
