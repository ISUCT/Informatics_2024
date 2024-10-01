package main

import (
	"fmt"
	"math"
)

func calc(x float64, b float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
}

func main() {
	fmt.Println(calc(1.28, 2.5))
	fmt.Println(calc(1.68, 2.5))
	fmt.Println(calc(2.08, 2.5))
	fmt.Println(calc(2.48, 2.5))
	fmt.Println(calc(2.88, 2.5))
	fmt.Println(calc(3.28, 2.5))
	fmt.Println(calc(1.1, 2.5))
	fmt.Println(calc(2.4, 2.5))
	fmt.Println(calc(3.6, 2.5))
	fmt.Println(calc(1.7, 2.5))
	fmt.Println(calc(3.9, 2.5))
}