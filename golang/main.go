package main

import (
	"fmt"

	"isuct.ru/informatics2022/laba4"
)

func main() {
	var y float64

	var a float64 = 2.5
	var b float64 = 4.6
	var x [5]float64 = [5]float64{1.2, 1.36, 1.57, 1.93, 2.25}

	for i := 0; i < len(x); i++ {
		y = laba4.Colcul(a, b, x[i])
		fmt.Println("при x =", x[i], " y =", y)
	}
}
