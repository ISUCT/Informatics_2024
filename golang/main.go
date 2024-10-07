package main

import (
	"isuct.ru/informatics2022/lab4"
)

const a float64 = 2.5

func main() {
	lab4.TaskA(a, 1.2, 3.7, 0.5)
	var s = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	lab4.TaskB(a, s)
}
