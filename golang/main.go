package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 0.1
	var b float64 = 0.5
	var exb = []float64{1.24, 1.38, 2.38, 3.21, 0.68}
	fmt.Println("Задача A")
	for x := 0.33; x <= 1.23; x += 0.18 {
		y := (math.Cbrt(a) + math.Pow(math.Tan(b*x), 4.5)) / (math.Pow(b, 0.2) + math.Pow((math.Cos(a*x)/math.Sin(a*x)), 2.7))
		fmt.Println(float32(y))
	}
	fmt.Println()
	fmt.Println("Задача B")
	for _, x := range exb {
		y := (math.Cbrt(a) + math.Pow(math.Tan(b*x), 4.5)) / (math.Pow(b, 0.2) + math.Pow((math.Cos(a*x)/math.Sin(a*x)), 2.7))
		fmt.Println(float32(y))
	}
}
