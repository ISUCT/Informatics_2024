package main

import (
	"fmt"
	"math"
)

func colcul(x float64) float64 {
	return math.Pow(math.Abs(float64(x*x-2.5)), 0.25) + math.Pow(math.Log10(x*x), 0.33333333)
}

func main() {
	var result float64
	var x [5]float64 = [5]float64{1.84, 2.71, 3.81, 4.56, 5.62}
	fmt.Println("Задача A")
	for i := 1.25; i <= 3.25; i += 0.40 {
		result = colcul(i)
		fmt.Println("x = ", i, "y = ", result)
	}
	fmt.Println("Задача B")
	for i := 0; i < len(x); i++ {
		result = colcul(x[i])
		fmt.Println("При x = ", x[i], "y = ", result)
	}
}
