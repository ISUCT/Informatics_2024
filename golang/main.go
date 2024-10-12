package main

import (
	"fmt"

	labs "isuct.ru/informatics2022/labs/lab4"
)

func main() {
	// Задание 4
	fmt.Println("------------------------------------------")
	fmt.Println(labs.Task_A(0.2, 2.2, 0.4))
	fmt.Println("------------------------------------------")
	arr := [5]float64{0.1, 0.9, 1.2, 1.5, 2.3}
	fmt.Println(labs.Task_B(arr[:]))
	fmt.Println("------------------------------------------")
}
