package lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculate_y(x float64) float64 {
	if math.Abs(x) >= 1 {
		return math.Pow(1.2, x) - math.Pow(x, 1.2)
	}
	return math.Acos(x)
}

// Задача под А
func Task_A(begin_x, end_x, delta_x float64) []float64 {
	var answer_arr []float64
	for x := begin_x; x < end_x; x += delta_x {
		answer_arr = append(answer_arr, calculate_y(x))
	}
	return answer_arr
}

// Задача под B
func Task_B(arguments []float64) []float64 {
	var answer_arr []float64
	for _, x := range arguments {
		answer_arr = append(answer_arr, calculate_y(x))
	}
	return answer_arr
}

func ReadDataFromTxt(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла %s: %v", filename, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var values []float64

	for fileScanner.Scan() {
		line := fileScanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка при конвертации строки '%s' в число: %v", line, err)
		}
		values = append(values, value)
	}

	return values, nil
}

func RunLab8ByLab4() {
	values, err := ReadDataFromTxt("labs/lab8/input.txt")
	if err != nil {
		panic(fmt.Sprintf("Ошибка при чтении данных: %v", err))
	}

	// Task_A
	begin_x := values[0]
	end_x := values[1]
	delta_x := values[2]
	fmt.Println("Task A")
	fmt.Println(Task_A(begin_x, end_x, delta_x))

	// Task_B
	arguments := values[3:8]
	fmt.Println("Task B")
	fmt.Println(Task_B(arguments))
}
