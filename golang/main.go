package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

const a float64 = 1.35
const b float64 = 0.98

func main() {
	A := lab4.TaskA(1.14, 4.24, 0.62)
	var B []float64 = []float64{0.35, 1.28, 3.51, 5.21, 4.16}

	fmt.Println("Задача A:")
	lab4.Enter(a, b, A)

	fmt.Println("Задача B:")
	lab4.Enter(a, b, B)
}
