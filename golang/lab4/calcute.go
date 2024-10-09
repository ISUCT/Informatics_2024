package main

import (
	"fmt"
	"math"
)

// Вариант 33
// Функция для вычисления значения y
func calculateY(x, a, b float64) float64 {
	numerator := 1 + math.Log2(x/a)
	denominator := b - x/math.Exp(a)
	return numerator / denominator
}

func main() {
	// Параметры функции
	a := 2.0
	b := 0.95

	// Задача A
	fmt.Println("Задача A:")
	xH := 1.25    // Начальное значение x
	xK := 2.75    // Конечное значение x
	deltaX := 0.3 // Шаг изменения x

	for x := xH; x <= xK; x += deltaX {
		y := calculateY(x, a, b)
		fmt.Printf("x = %.2f, y = %.4f\n", x, y)
	}

	// Задача B
	fmt.Println("Задача B:")
	xValues := []float64{2.2, 3.78, 4.51, 6.58, 1.2}

	for _, x := range xValues {
		y := calculateY(x, a, b)
		fmt.Printf("x = %.2f, y = %.4f\n", x, y)
	}
}
