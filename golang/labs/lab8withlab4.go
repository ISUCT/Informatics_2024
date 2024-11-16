package labs

import (
	"bufio"
	"fmt"

	"math"
	"os"
	"strconv"
)

func readDataFromFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var floats []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга float '%s': %w", line, err)
		}
		floats = append(floats, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return floats, nil
}

func CalculateY(x float64, a float64, b float64) float64 {

	y := math.Acos(math.Pow(x, 2)-math.Pow(b, 2)) / math.Asin(math.Pow(x, 2)-math.Pow(a, 2))
	return y
}

func TaskA(a, b, xn, xk, deltax float64) [][]float64 {
	var taskA [][]float64
	for x := xn; x <= xk; x += deltax {
		y := CalculateY(x, a, b)
		taskA = append(taskA, []float64{x, y})
	}
	return taskA
}

func TaskB(a, b float64, values []float64) [][]float64 {
	var taskB [][]float64
	for _, x := range values {
		y := CalculateY(x, a, b)
		taskB = append(taskB, []float64{x, y})
	}
	return taskB
}

func RunLab4() {
	data, err := readDataFromFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(data) < 5 {
		fmt.Println("Недостаточно значений в файле! Попробуй еще раз")
		return
	}
	a := data[0]
	b := data[1]
	xn := data[2]
	xk := data[3]
	deltax := data[4]
	otherValues := data[5:]

	resultsA := TaskA(a, b, xn, xk, deltax)
	for _, result := range resultsA {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}

	arr := otherValues
	resultsB := TaskB(a, b, arr)
	for _, result := range resultsB {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}
}
