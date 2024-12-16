package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(fileName string) (string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return string(file), nil
}

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

	if len(variables) < 5 {
		return variables, fmt.Errorf("нехватает значений! попробуй еще раз")
	}

	return variables, fileScanner.Err()
}
