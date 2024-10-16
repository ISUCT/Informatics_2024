package main

import (
	"fmt"

	fourlaba "isuct.ru/informatics2022/4lab"
)

func solveTaskA(a, b, xStart, xEnd, step float64) {
	for x := xStart; x <= xEnd; x += step {
		y := fourlaba.CalculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.5f\n", x, y)
	}
}

func solveTaskB(a, b float64, xValues []float64) {
	for _, x := range xValues {
		y := fourlaba.CalculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.5f\n", x, y)
	}
}

func main() {
	fmt.Println("Семенов Алексей Дмитриевич")
	a := 2.0
	b := 4.1

	fmt.Println("Задача A:")
	xStartA := 0.77
	xEndA := 1.77
	stepA := 0.2
	solveTaskA(a, b, xStartA, xEndA, stepA)

	fmt.Println("Задача B:")
	xValuesB := []float64{1.24, 1.38, 2.38, 3.21, 0.68}
	solveTaskB(a, b, xValuesB)
}
