package main

import (
	"fmt"
	"isuct.ru/informatics2022/laba4"
)

func main() {
	var a float64 = 0.4
	var b float64 = 0.8
	var y float64
	var x [5] float64 = [5] float64 {4.48, 3.56, 2.78, 5.28, 3.21}
	// Задание A
	for i := 3.2; i<3.21; i+=0.6{
		y = laba4.Lab4(a,b,i)
		fmt.Println(y)
	}
	// Задание B
	for i := 0; i<len(x)-1; i++{
		y = laba4.Lab4(a,b,x[i])
		fmt.Println(y)
	}
}

