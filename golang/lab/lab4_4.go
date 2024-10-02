package main

import (
	"fmt"
	"math"
)

const b float64 = 2.5

func calc(b float64, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
}

func taskA(b float64, Xn float64, Xk float64, delX float64) {
	for x := Xn; x <= Xk; x += delX {
		fmt.Println(calc(b, x))
	}
}

func taskB(b float64, x [5]float64) {
	for _, value := range x {
		fmt.Println(calc(b, value))
	}
}

func main() {
	taskA(b, 1.28, 3.28, 0.4)
	var slice = [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
	taskB(b, slice)
}
