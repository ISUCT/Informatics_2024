package main

import (
	"fmt"

	"isuct.ru/informatics2022/laba4"
)

func main() {
	var result float64

	var a float64 = 2.0
	var b float64 = 4.1
	var x [5]float64 = [5]float64{1.24, 1.38, 2.38, 3.21, 0.68}

	fmt.Println("задача А:")
	for i := 0.77; i < 1.77; i += 0.2 {
		result = laba4.Colcul(a, b, i)
		fmt.Printf("при x = %.2f, y = %.4f \n", i, result)
	}

	fmt.Println("задача B:")
	for i := 0; i < len(x); i++ {
		result = laba4.Colcul(a, b, x[i])
		fmt.Printf("при x = %.2f, y = %.4f \n", x[i], result)
	}
}
