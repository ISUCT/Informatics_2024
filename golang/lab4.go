package main

import (
	"log"
	"math"
)

// Функция для вычисления значения y по формуле
func lab4(x, a, b float64) float64 {
	if x <= 1 {
		log.Fatalf("Ошибка: x должно быть больше 1 для корректных вычислений.")
	}
	return (math.Pow(a, 3)*math.Sqrt(x) - b*math.Log(x)) / (math.Pow(math.Log10(x-1), 3))
}
