package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab"
)

func main() {
	// Лабораторная №3
	fmt.Println("Белов Дмитрий Алексеевич")

	// Лабораторная №4
	const b float64 = 2.5
	var slice = [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
	ValuesA := lab.TaskA(b, 1.28, 3.28, 0.4)
	ValuesB := lab.TaskB(b, slice)
	lab.RunLab4Tasks(ValuesA)
	lab.RunLab4Tasks(ValuesB)
}
