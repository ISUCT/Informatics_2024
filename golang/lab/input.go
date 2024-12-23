package lab

import (
	"bufio"
	"fmt"
	"os"
)

func InputValue() string {
	inputs := ""

	file_read := bufio.NewReader(os.Stdin)
	fmt.Println("Введите данные:")

	for i := 0; i < 10; i++ {
		input, err := file_read.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(0)
		}
		inputs += input
	}
	return inputs
}
