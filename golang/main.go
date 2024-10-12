package main

import (
	"fmt"

	"isuct.ru/informatics2022/fourlaba"
)

func main() {
	a := 2.0
	b := 4.1

	// Задача A: вычисляем y для значений x от 0.77 до 1.77 с шагом 0.2
	fmt.Println("Задача A:")
	xStartA := 0.77
	xEndA := 1.77
	step := 0.2

	for x := xStartA; x <= xEndA; x += step {
		y := fourlaba.CalculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.5f\n", x, y)
	}

	fmt.Println("Задача B:")
	xValuesB := []float64{1.24, 1.38, 2.38, 3.21, 0.68}

	// Цикл для задачи B
	for _, x := range xValuesB {
		y := fourlaba.CalculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.5f\n", x, y)
	}
}
