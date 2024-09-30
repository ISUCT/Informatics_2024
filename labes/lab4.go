package main

import (
	"fmt"
	"math"
)

// Функция для решения дифференциального уравнения y' = f(x).
func EulerMethod(f func(float64) float64, x0, h float64, n int) []float64 {
	var xs, ys []float64
	xs = append(xs, x0)
	ys = append(ys, 0.0)

	for i := 1; i <= n; i++ {
		yNew := ys[i-1] + h*f(xs[i-1])
		xs = append(xs, xs[i-1]+h)
		ys = append(ys, yNew)
	}

	return ys
}

// Функция f(x) для данного дифференциального уравнения.
func f(x float64) float64 {
	return 4.1*math.Pow(math.Sqrt(x), 3) - 2.7*math.Log(5)*x
}

func main() {
	const a = 4.1 // Параметры уравнения
	const b = 2.7
	const x0 = 1.9 // Начальное значение x
	const h = 0.4  // Шаг по x
	const n = 10   // Количество шагов

	// Решение уравнения методом Эйлера
	result := EulerMethod(f, x0, h, n)

	// Вывод результата
	fmt.Println("Результаты расчета")
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d\t%.2f\n", i+1, result[i])
	}
}
