package main

import (
	"fmt"
	"math"
)

func laba4() {
	// Значения параметров
	a := 1.3
	b := 2.4
	x := 1.88

	// Вычисление функции y = sqrt(x^2 - (a * x + b)^2)
	y := math.Sqrt(math.Pow(x, 2) - math.Pow(a*x+b, 2))

	// Вывод результата
	fmt.Printf("Значение y для x=%.2f: %.4f\n", x, y)
}
