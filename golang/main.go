package main

import (
	"fmt"
	"isuct.ru/informatics2022/lab4"
)

func main() {
	var a float64 = 7.2
	var b float64 = 1.3
	var x []float64 = []float64{2.4, 2.8, 3.9, 4.7, 3.16}
	var xn float64 = 1.56
	var xk float64 = 4.71
	var xdel float64 = 0.63

	var resA []float64 = lab4.TaskA(a, b, xn, xk, xdel)
	fmt.Println("Задача А", resA)
	var resB []float64 = lab4.TaskB(a, b, x)
	fmt.Println("Задача В", resB)
	fmt.Println("Нерабеев Кирилл Сергеевич")
}
