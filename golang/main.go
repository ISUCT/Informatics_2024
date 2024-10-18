package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	var result float64
	var a float64 = 1.1
	var b float64 = 0.09
	var x [5]float64 = [5]float64{1.21, 1.76, 2.53, 3.48, 4.52}
	fmt.Println("\nзадача A:")
	for i := 1.2; i < 2.2; i += 0.2 {
		result = lab4.Colcul(a, b, i)
		fmt.Println("при x=", i, "y=", result)
	}
	fmt.Println("\nзадача B:")
	for _, i := range x {
		result = lab4.Colcul(a, b, i)
		fmt.Println("при x=", i, "y=", result)
	}
}
