package main

import (
	"fmt"
)

func main() {
	var a float64 = 4.1
	var b float64 = 2.7

	// Срезы для хранения значений x и y
	var xValuesA, yValuesA []float64
	var xValuesB, yValuesB []float64

	// Задача А
	fmt.Println("Задача А:")
	xStart := 1.5
	xEnd := 3.5
	xDelta := 0.4

	// Заполняем срезы значениями x и y для задачи А
	for x := xStart; x <= xEnd; x += xDelta {
		xValuesA = append(xValuesA, x)
		yValuesA = append(yValuesA, lab4(x, a, b))
	}

	// Выводим результаты задачи А
	for i := range xValuesA {
		fmt.Printf("x = %.2f, y = %.5f\n", xValuesA[i], yValuesA[i])
	}

	// Задача Б
	fmt.Println("\nЗадача Б:")
	xValuesB = []float64{1.9, 2.15, 2.34, 2.74, 3.16}

	// Заполняем срез y для задачи Б
	for _, x := range xValuesB {
		yValuesB = append(yValuesB, lab4(x, a, b))
	}

	// Выводим результаты задачи Б
	for i := range xValuesB {
		fmt.Printf("x = %.2f, y = %.5f\n", xValuesB[i], yValuesB[i])
	}
}
