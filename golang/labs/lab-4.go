package main

import (
	"fmt"
	"math"
)

func logarifm(x, base float64) float64 {
	return math.Log(x) / math.Log(base)
}

func Calc(a, b, x float64) float64 {
	koren := math.Cbrt(x)
	log5 := logarifm(x, 5)
	lgkub := math.Pow(math.Log(x-1), 3)
	y := (a*koren - b*log5) / lgkub
	return y
}

func main() {
	var a float64 = 4.1
	var b float64 = 2.7

	fmt.Println("Задача A:")
	for x := 1.5; x <= 3.5; x += 0.4 {
		y := Calc(a, b, x)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}

	fmt.Println("\nЗадача B:")
	xZnach := []float64{1.9, 2.15, 2.34, 2.74, 3.16}
	for _, x := range xZnach {
		y := Calc(a, b, x)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}
}
