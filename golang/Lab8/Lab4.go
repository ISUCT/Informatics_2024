package Lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calcYA(b, x float64) float64 {
	denominator := math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3))
	if denominator == 0 {
		return math.Inf(1)
	}
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / denominator
}

func taskAA(bA, xStart, xEnd, deltaX float64) []float64 {
	var results []float64
	for x := xStart; x <= xEnd; x += deltaX {
		y := calcYA(bA, x)
		results = append(results, y)
	}
	return results
}

func taskBA(bB float64, xValues []float64) []float64 {
	var results []float64
	for _, x := range xValues {
		y := calcYA(bB, x)
		results = append(results, y)
	}
	return results
}

func RunLab8A() {
	const filename = "input.txt"

	values, err := readInputFileA(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	if len(values) < 2 {
		fmt.Println("Недостаточно значений в файле")
		return
	}

	bA := values[0]
	bB := values[1]
	xValues := values[2:]

	xStart := 1.28
	xEnd := 3.28
	deltaX := 0.4

	resultsA := taskAA(bA, xStart, xEnd, deltaX)
	fmt.Println("Результаты задания A:")
	for i, y := range resultsA {
		x := xStart + float64(i)*deltaX
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}

	resultsB := taskBA(bB, xValues)
	fmt.Println("\nРезультаты задания B:")
	for i, y := range resultsB {
		x := xValues[i]
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}
}

func readInputFileA(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var values []float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Ошибка при чтении числа:", err)
			continue
		}
		values = append(values, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return values, nil
}
