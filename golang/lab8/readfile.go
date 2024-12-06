package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileForLab4(fileName string) []float64 {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Не удается открыть файл")
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var variables []float64

	for fileScanner.Scan() {
		number, err := strconv.ParseFloat(fileScanner.Text(), 64)
		if err != nil {
			fmt.Println("Ошибка при переводе числа")
			panic(err)
		}
		variables = append(variables, number)
	}
	return variables
}
