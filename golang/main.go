package main

import (
	"fmt"

	"isuct.ru/informatics2022/code"
)

func main() {
	var a float64 = 2.0
	var xs = [8]float64{1.2, 4.2, 0.6, 1.16, 1.32, 1.47, 1.65, 1.93}
	for _, x := range xs {
		fmt.Println("y =", func_.Lab_4(a, x))
	}
}
