package main

import (
	"isuct.ru/informatics2022/lab4"
	"fmt"
)

const a float64 = 1.1
const b float64 = 0.09

func main() {
	fmt.Println("Task A:")
	lab4.TaskA(a,b,1.2,2.2,0.2)
	fmt.Println("Task B:")
	var s = [5]float64{1.21,1.76,2.53,3.48,4.52}
	lab4.TaskB(a, b, s)
}
