package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	var a float64 = 7.2
	var b float64 = 4.2
	return (math.Sqrt(math.Abs(a-b*x)) / math.Pow(math.Log(x), 3))
}

func main() {
	var exb = []float64{2.4, 2.8, 3.9, 4.7, 3.16}
	fmt.Println("Задача A")
	for x := 1.81; x <= 5.31; x += 0.7 {
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
