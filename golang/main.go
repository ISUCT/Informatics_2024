package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"

	"isuct.ru/informatics2022/lab6"
)

func main() {
	fmt.Println("Губе Владислав Семёнович")
	fmt.Println("Task A:\n", lab4.Task_A(3.2, 6.2, 0.6))
	fmt.Println("Task B:\n", lab4.Task_B([]float64{4.48, 3.56, 2.78, 5.28, 3.21}))
	lab6.Lab6()
}
