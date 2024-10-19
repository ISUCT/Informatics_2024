package lab4

import (
	"math"

	"fmt"
)

func uravnenie (b, x float64) float64 {
	var b3 float64 = math.Pow(b, 3)
	var x3 float64 = math.Pow(x, 3)

	return (1+math.Pow(math.Sin(b3 + x3), 2))/math.Pow(b3 + x3, 1/3)
}


func for_A (b, xbegin, xend, xdelta float64) []float64 {
	var ys []float64
	for x := xbegin; x < xend; x += xdelta {
		ys = append(ys, uravnenie(b, x))
	}
	return ys
}

func for_B (b float64, xs []float64) []float64 {
	var ys []float64
	for _, x := range xs {
		ys = append(ys, uravnenie(b, x))
	}
	return ys
}

func Ans_lab4 () {
	var b float64 = 2.5
	var xbegin float64 = 1.28
	var xend float64 = 3.28
	var xdelta float64 = 0.4
	var xs []float64 = []float64{1.1, 2.4, 3.6, 1.7, 3.9}
	var for_a = for_A(b, xbegin, xend, xdelta)
	var for_b = for_B(b, xs)

	fmt.Println("Решение задания А:")
	for i, y := range for_a {
		fmt.Printf("y%d = %f\n", i+1, y)
	}

	fmt.Println("Решение задания B:")
	for i, y := range for_b {
		fmt.Printf("y%d = %f\n", i+1, y)
	}
}