package main

import (
	"fmt"
	"math"
)

func main() {
	// Задача A: Найти значения функции y для xH - xK с шагом x_delta
	xH := 0.22
	xK := 0.92
	xDelta := 0.14

	// Функция для вычисления y
	yFunc := func(x float64) float64 {
		innerFunc := math.Pow(math.Asin(x), 4) + math.Pow(math.Acos(x), 6)
		return math.Pow(innerFunc, 1.0/7.0)
	}

	var results []float64

	for x := xH; x <= xK; x += xDelta {
		results = append(results, yFunc(x))
	}

	fmt.Println("Значения y для x от", xH, "до", xK, "с шагом", xDelta)
	fmt.Println("Результаты:")
	for _, result := range results {
		fmt.Printf("%v\n", result)
	}

	// Задача B: Найти значения функции для заданных x
	xs := []float64{0.1, 0.35, 0.4, 0.55, 0.6}

	fmt.Println("Значения y для x от", xs[0], "до", xs[len(xs)-1])
	fmt.Println("Результаты:")
	for _, x := range xs {
		res := yFunc(x)
		fmt.Printf("%v\n", res)
	}
}
