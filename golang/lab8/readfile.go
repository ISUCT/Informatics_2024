package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileForLab4(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("не удается открыть файл")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var variables []float64

	for fileScanner.Scan() {
		number, err := strconv.ParseFloat(fileScanner.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка при переводе числа")
		}
		variables = append(variables, number)
	}

	fmt.Printf("Содержание файла %v \n", variables)
	return variables, fileScanner.Err()
}
