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
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var floats []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing float '%s': %w", line, err)
		}
		floats = append(floats, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return floats, nil
}

func CalculateY(x, a, b float64) (float64, error) {
	argAcos := math.Pow(x, 2) - math.Pow(b, 2)
	argAsin := math.Pow(x, 2) - math.Pow(a, 2)

	if argAcos < -1 || argAcos > 1 {
		return 0, fmt.Errorf("math.Acos argument out of range: %f", argAcos)
	}
	if argAsin < -1 || argAsin > 1 {
		return 0, fmt.Errorf("math.Asin argument out of range: %f", argAsin)
	}
	if argAsin == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	y := math.Acos(argAcos) / math.Asin(argAsin)
	return y, nil
}

func TaskA(a, b, xn, xk, deltax float64) ([][]float64, error) {
	var taskA [][]float64
	for x := xn; x <= xk; x += deltax {
		y, err := CalculateY(x, a, b)
		if err != nil {
			return nil, fmt.Errorf("error calculating y for x=%f: %w", x, err)
		}
		taskA = append(taskA, []float64{x, y})
	}
	return taskA, nil
}

func TaskB(a, b float64, values []float64) ([][]float64, error) {
	var taskB [][]float64
	for _, x := range values {
		y, err := CalculateY(x, a, b)
		if err != nil {
			return nil, fmt.Errorf("error calculating y for x=%f: %w", x, err)
		}
		taskB = append(taskB, []float64{x, y})
	}
	return taskB, nil
}

func RunLab4() {
	data, err := readDataFromFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(data) < 5 {
		fmt.Println("Not enough values in the file! Try again")
		return
	}
	a := data[0]
	b := data[1]
	xn := data[2]
	xk := data[3]
	deltax := data[4]
	otherValues := data[5:]

	resultsA, err := TaskA(a, b, xn, xk, deltax)
	if err != nil {
		fmt.Println("Error in TaskA:", err)
		return
	}
	for _, result := range resultsA {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}

	resultsB, err := TaskB(a, b, otherValues)
	if err != nil {
		fmt.Println("Error in TaskB:", err)
		return
	}
	for _, result := range resultsB {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}
}
