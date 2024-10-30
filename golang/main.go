package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Println("Коротков Денис Викторович")
	fmt.Println("Задача А:\n", lab4.TaskA(0.26, 0.66, 0.08))
	fmt.Println("Задача В:\n", lab4.TaskB([]float64{0.1, 0.35, 0.4, 0.55, 0.6}))
}
