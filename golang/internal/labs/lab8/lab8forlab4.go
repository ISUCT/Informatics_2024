package lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
			return nil, fmt.Errorf("че не открывается? иди делом займись")
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func Calculate(x float64) float64 {
	var n float64 = 1.0 / 7
	return math.Pow(math.Pow(math.Asin(x), 4)+math.Pow(math.Acos(x), 6), n)
}

func Task_A(beginx float64, endx float64, stepx float64) []float64 {
	answers := []float64{}
	for x := beginx; x <= endx; x += stepx {
		result := Calculate(x)
		answers = append(answers, result)
	}
	return answers
}

func Task_B(arr []float64) []float64 {
	answers := []float64{}
	for _, x := range arr {
		answers = append(answers, Calculate(x))
	}
	return answers
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
	for i, value := range Task_A(beginx, endx, stepx) {
		fmt.Printf("x = %d\ty = %f\n", i, value)
	}
	fmt.Println("Задание B")
	for i, value := range Task_B(taskBslice) {
		fmt.Printf("x = %d\ty = %f\n", i, value)
	}
	return nil
}
