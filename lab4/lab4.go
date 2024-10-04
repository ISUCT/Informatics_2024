package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	var a float64 = 0.1
	var b float64 = 0.5
	return (math.Cbrt(a) + math.Pow(math.Tan(b*x), 4.5)) / (math.Pow(b, 0.2) + (1 / (math.Pow(math.Tan(a*x), 2.7))))
}

func main() {
	var exb = []float64{1.24, 1.38, 2.38, 3.21, 0.68}
	fmt.Println("Задача A")
	for x := 0.33; x <= 1.23; x += 0.18 {
		y := f(x)
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
	fmt.Println()
	fmt.Println("Задача B")
	for _, x := range exb {
		y := f(x)
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
}
