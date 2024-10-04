package main

import (
	"fmt"
	"math"
)

func function(x float64) {
	var b float64 = 2.5
	y := (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt((math.Pow(b, 3) + math.Pow(x, 3))))
	fmt.Println(y)
}

func main() {
	fmt.Print("Задача А\n")
	var x_a float64 = 1.28
	for i := 0; i < 6; i++ {
		function(x_a)
		x_a = x_a + 0.4
	}
	fmt.Print("Задача В\n")
	xs := [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
	for _, x_b := range xs {
		function(x_b)
	}
}
