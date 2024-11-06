package lab4

import (
	"fmt"
	"math"
)

func Lab4A(xn, xk, dx float64) float64 {
	fmt.Println("\nЗадача А")
	for x := xn; x <= xk; x += dx {
		fmt.Printf("y(%f) = %f\n", x, Lab4Y(x))
	}
	return 0
}

func Lab4B(x1, x2, x3, x4, x5 float64) float64 {
	fmt.Println("Задача B")
	fmt.Printf("y(%f) = %f\n", x1, Lab4Y(x1))
	fmt.Printf("y(%f) = %f\n", x2, Lab4Y(x2))
	fmt.Printf("y(%f) = %f\n", x3, Lab4Y(x3))
	fmt.Printf("y(%f) = %f\n", x4, Lab4Y(x4))
	fmt.Printf("y(%f) = %f\n", x5, Lab4Y(x5))
	return 0
}

func Lab4Y(x float64) float64 {
	return (math.Pow(math.Pow(math.Asin(x), 4)+math.Pow(math.Acos(x), 6), 1.0/7.0))
}

func Lab4() {
	var xn, xk, dx float64 = 0.22, 0.92, 0.14
	var x1, x2, x3, x4, x5 float64 = 0.1, 0.35, 0.4, 0.55, 0.6
	Lab4A(xn, xk, dx)
	Lab4B(x1, x2, x3, x4, x5)
}
