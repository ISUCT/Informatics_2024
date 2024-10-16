package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Println("Трофимов Илья Михайлович")
	resultA := lab4.TaskA(2.5, 4.6, 1.15, 3.05, 0.38)
	fmt.Println(resultA)
	resultB := lab4.TaskB(2.5, 4.6, []float64{1.2, 1.36, 1.57, 1.93, 2.25})
	fmt.Println(resultB)
}
