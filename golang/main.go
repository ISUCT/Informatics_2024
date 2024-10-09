package main

import (
	"isuct.ru/informatics2022/lab"
	"fmt"
)

const b float64 = 2.5

func main() {
	var slice = [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
	ValuesA := lab.TaskA(b, 1.28, 3.28, 0.4)
	ValuesB := lab.TaskB(b, slice)
	for _, value := range ValuesA {
		fmt.Println(value)
	}
	for _, value := range ValuesB {
		fmt.Println(value)
	}
}