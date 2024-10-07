package main

import "isuct.ru/informatics2022/lab4"

const a float64 = 1.35
const b float64 = 0.98

func main() {
	lab4.TaskA(b, a, 1.14, 4.24, 0.62)
	var s = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
	lab4.TaskB(b, a, s)
}
