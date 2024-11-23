package labs

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ReadInput(filename string) (a, b float64, xValues []float64, err error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, 0, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return 0, 0, nil, fmt.Errorf("ошибка преобразования строки в число: %w", err)
		}
		if i == 0 {
			a = num
		} else if i == 1 {
			b = num
		} else {
			xValues = append(xValues, num)
		}
		i++
	}
	return a, b, xValues, nil
}

func RunLab4() {
	a, b, xValues, err := ReadInput("input.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println(TaskA(xValues[0], xValues[1], xValues[2], a, b))
	fmt.Println(TaskB(xValues[3:], a, b))
}

func TaskA(xn, xk, deltax, a, b float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += deltax {
		yValues = append(yValues, Calculate_y(x, a, b))
	}
	return yValues
}

func TaskB(values []float64, a, b float64) []float64 {
	var yValues []float64
	for _, x := range values {
		yValues = append(yValues, Calculate_y(x, a, b))
	}
	return yValues
}

func Calculate_y(x float64, a float64, b float64) float64 {
	y := math.Pow(a+b*x, 2.5) / (1 + math.Log10(a+b*x))
	return y
}
