package main

import (
	"fmt"

	"isuct.ru/informatics2022/laba4"
)

func main() {
	var a float64 = 2.5
	var b float64 = 4.6
	var xMin float64 = 1.15
	var xMax float64 = 3.05
	var xDelta float64 = 0.38
	var x []float64 = []float64{1.2, 1.36, 1.57, 1.93, 2.25}

	fmt.Println("задача А")
	laba4.TaskA(a, b, xMin, xMax, xDelta)

	fmt.Println("задача B")
	laba4.TaskB(a, b, x)
}