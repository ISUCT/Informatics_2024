package lab4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func RunLab4() {
	a, b, values := ReadInputFile("input.txt")

	fmt.Println(TaskA(1.1, 3.6, 0.5, a, b))
	fmt.Println(TaskB(values, a, b))
}

func ReadInputFile(filename string) (float64, float64, []float64) {
	var a, b float64
	var values []float64

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return 0, 0, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Ошибка чтения числа:", err)
			continue
		}

		if i == 0 {
			a = num
		} else if i == 1 {
			b = num
		} else {
			values = append(values, num)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка сканирования файла:", err)
	}

	return a, b, values
}

func TaskA(xn, xk, deltax, a, b float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += deltax {
		yValues = append(yValues, CalculateY(x, a, b))
	}
	return yValues
}

func TaskB(values []float64, a, b float64) []float64 {
	var yValues []float64
	for _, x := range values {
		yValues = append(yValues, CalculateY(x, a, b))
	}
	return yValues
}

func CalculateY(x float64, a float64, b float64) float64 {
	y := math.Pow(a+b*x, 2.5) / (1 + math.Log10(a+b*x))
	return y
}
