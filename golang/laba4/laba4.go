package laba4

import (
	"fmt"
	"math"
)

// Функция для вычисления значения
func calculate(x float64) float64 {
	return math.Pow(1.2, x) - math.Pow(x, 1.2)
}

// Функция по задаче 6
func Y_A() {
	var begin_x float64 = 0.2
	var end_x float64 = 2.2
	var delta_x float64 = 0.4

	for begin_x < end_x {
		x := begin_x
		answer := calculate(x) // Используем новую функцию
		fmt.Println(answer, x)

		begin_x += delta_x
	}
}

// Задача 6.B
func Y_B() {
	var arr = [5]float64{0.1, 0.9, 1.2, 1.5, 2.3}
	for _, x := range arr {
		answer := calculate(x) // Используем новую функцию
		fmt.Println(answer)
	}
}
