package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	const a float64 = 2.0
	const b float64 = 0.95
	var p = [5]float64{2.2, 3.78, 4.51, 6.58, 1.2}

	fmt.Println("Щербаков Илья Алексадрович")
	fmt.Println("Task А:\n", lab4.TaskA(b, a, 1.25, 2.75, 0.3))
	fmt.Println("Task В:\n", lab4.TaskB(a, b, p))
}
