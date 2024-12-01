package labs

import (
	"fmt"
	"math"

	lab8 "isuct.ru/informatics2022/labs/package_lab8"
)

func RunLab4() {
	a, b, xValues, err := lab8.ReadInput("input.txt")
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
