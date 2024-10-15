package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

const a float64 = 1.35
const b float64 = 0.98

func main() {
	var A [6]float64 = [6]float64{1.14, 1.76, 2.38, 3, 3.62, 4.24}
	var B [5]float64 = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}

	fmt.Println("Задача A:")
	for i := 0; i < 6; i++ {
		y := lab4.Calculator(a, b, A[i])
		fmt.Printf("X%d = %g\n", i, y)
	}

	fmt.Println("Задача B:")
	for i := 0; i < 5; i++ {
		y := lab4.Calculator(a, b, B[i])
		fmt.Printf("X%d = %g\n", i, y)
	}
}
