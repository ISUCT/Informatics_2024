package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"isuct.ru/informatics2022/labs/lab4"
)

func ReadDataFromTxt(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла %s: %w", filename, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var values []float64

	for fileScanner.Scan() {
		line := fileScanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка при конвертации строки '%s' в число: %w", line, err)
		}
		values = append(values, value)
	}

	return values, nil
}

func RunLab8ByLab4() {
	values, err := ReadDataFromTxt("labs/lab8/input.txt")
	if err != nil {
		panic("Ошибка при чтении данных")
	}

	// Task_A
	begin_x := values[0]
	end_x := values[1]
	delta_x := values[2]
	fmt.Println("Task A")
	fmt.Println(lab4.Task_A(begin_x, end_x, delta_x))

	// Task_B
	arguments := values[3:8]
	fmt.Println("Task B")
	fmt.Println(lab4.Task_B(arguments))
}
