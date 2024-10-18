package main

import (
	"fmt"
	"lab4/calculate"
)

func main() {
	var a, b, xStart, xEnd, step float64

	a = 0.1
	b = 0.5
	xStart = 0.33
	xEnd = 1.23
	step = 0.18
	fmt.Printf("Task A:\n")
	for x := xStart; x <= xEnd; x += step {
		y := calculate.CalculateY(a, b, x)
		fmt.Printf("For x = %.2f, y = %.10f\n", x, y)
	}

	fmt.Printf("Task B:\n")
	x1 := 0.5
	x2 := 0.36
	x3 := 0.40
	x4 := 0.62
	x5 := 0.78

	fmt.Printf("For x1 = %.2f, y = %.10f\n", x1, calculate.CalculateY(a, b, x1))
	fmt.Printf("For x2 = %.2f, y = %.10f\n", x2, calculate.CalculateY(a, b, x2))
	fmt.Printf("For x3 = %.2f, y = %.10f\n", x3, calculate.CalculateY(a, b, x3))
	fmt.Printf("For x4 = %.2f, y = %.10f\n", x4, calculate.CalculateY(a, b, x4))
	fmt.Printf("For x5 = %.2f, y = %.10f\n", x5, calculate.CalculateY(a, b, x5))
}
