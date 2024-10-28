package lab4

import (
	"fmt"
	"math"
)

func yFunc(x float64) float64 {
	return math.Pow(math.Sqrt(math.Asin(x)*math.Asin(x)*math.Asin(x)*math.Asin(x)+math.Acos(x)*math.Acos(x)*math.Acos(x)*math.Acos(x)), 1.0/7.0)
}

func Laba() {
	// Задача A: Найти значения функции y для xH - xK с шагом x_delta
	xH := 0.22
	xK := 0.92
	x_delta := 0.14
	var results []float64

	for x := xH; x <= xK; x += x_delta {
		results = append(results, yFunc(x))
	}

	fmt.Println("Значения y для x от", xH, "до", xK, "c шагом", x_delta)
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
