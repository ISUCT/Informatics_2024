package main

import (
	"fmt"
	"math"
)

func calc(x float64, b float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
}

// в цикле на второй итерации выдается огромное число: 1.64000000000000001
func main() {
	for i := 1.28; i <= 3.28; i += 0.4 {
		fmt.Println(calc(i, 2.5))
	}
	fmt.Println(calc(1.1, 2.5))
	fmt.Println(calc(2.4, 2.5))
	fmt.Println(calc(3.6, 2.5))
	fmt.Println(calc(1.7, 2.5))
	fmt.Println(calc(3.9, 2.5))
}
