package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	var (
		a float64 = 1.6
		y float64
	)
	var x [5]float64 = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	for i := 0; i < len(x)-1; i++ {
		y = lab4.Lab4(a, x[i])
		fmt.Println(y)
	}
	for i := 1.2; i < 3.7; i += 0.5 {
		y = lab4.Lab4(a, i)
		fmt.Println(y)
	}
}
