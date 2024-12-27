package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"isuct.ru/informatics2022/internal/labs/lab4"
)

func ReadData(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var numbers []float64
	for sc.Scan() {
		number, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("не открывается")
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func RunLab8forLab4(filename string) error {
	values, err := ReadData(filename)
	if err != nil {
		return err
	}
	var beginx float64 = values[0]
	var endx float64 = values[1]
	var stepx float64 = values[2]
	var taskBslice = values[3:]
	fmt.Println("Задача А")
	for i, value := range lab4.Task_A(beginx, endx, stepx) {
		fmt.Printf("x = %d\ty = %f\n", i, value)
	}
	fmt.Println("Задание B")
	for i, value := range lab4.Task_B(taskBslice) {
		fmt.Printf("x = %d\ty = %f\n", i, value)
	}
	return nil
}
