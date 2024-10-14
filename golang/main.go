package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

const b float64 = 2.5

func main() {
	fmt.Println(lab4.TaskA(b, 1.2, 3.7, 0.5))
	var s = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	fmt.Println(lab4.TaskB(b, s))
}
