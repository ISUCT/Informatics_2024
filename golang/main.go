package main

import (
	"isuct.ru/informatics2022/lab4"
)


func main() {
	const a float64 = 2.0
	const b float64 = 0.95
	lab4.F_A(b, a, 1.25, 2.75, 0.3)
	var p = [5]float64{2.2, 3.78, 4.51, 6.58, 1.2}
	lab4.F_B(a, b, p)
}