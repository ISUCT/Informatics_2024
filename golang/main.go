package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	var a float64 = 2.0
	var b float64 = 0.95
	var xb float64 = 1.25
	var xe float64 = 2.75
	var xd float64 = 0.3

	var x []float64 = []float64{2.2, 3.78, 4.51, 6.58, 1.2}

	fmt.Print("task A ")
	fmt.Println(lab4.A(a, b, xb, xe, xd))
	fmt.Print("task B ")
	fmt.Println(lab4.B(a, b, x))

	fmt.Println("Шумская Дарья")
}
