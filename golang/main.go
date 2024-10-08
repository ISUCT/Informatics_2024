package main

import (
	"isuct.ru/informatics2022/lab4"
)

const a float64 = 2.0

func main() {
	lab4.TaskA(a, 1.08, 1.88, 0.16)
	var slice = [5]float64{1.16, 1.35, 1.48, 1.52, 1.96}
	lab4.TaskB(a, slice)
}
