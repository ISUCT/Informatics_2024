package main

import (
	"isuct.ru/informatics2022/lab"
)

const b float64 = 2.5

func main() {
	TaskA(b, 1.28, 3.28, 0.4)
	var slice = [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
	TaskB(b, slice)
}