package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

const a float64 = 1.35
const b float64 = 0.98

func main() {
	var A []float64 = []float64{1.14, 1.76, 2.38, 3, 3.62, 4.24}
	var B []float64 = []float64{0.35, 1.28, 3.51, 5.21, 4.16}

	fmt.Println("Задача A:")
	lab4.Enter(a, b, A)

	fmt.Println("Задача B:")
	lab4.Enter(a, b, B)
}
