package main

import (
	"fmt"
	"math"
)

func calculate(x float64) float64 {
	var k float64 = 1
	var l float64 = 7
	var n float64 = k / l
	return math.Pow(math.Pow(math.Asin(x), 4)+math.Pow(math.Acos(x), 6), n)
}

func main() {

	for x := 0.22; x <= 0.92; x += 0.14 {
		result := calculate(x)
		fmt.Printf("результат от %.2f : %.2f\n", x, result)

	}
}
