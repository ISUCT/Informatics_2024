package main

import (
	"fmt"

	"isuct.ru/informatics2022/laba4"
)

func main() {
	var result float64

	var a float64 = 2.5
	var b float64 = 4.6
	var x [5]float64 = [5]float64{1.2, 1.36, 1.57, 1.93, 2.25}

	fmt.Println("\nзадача А:")
	for i := 1.15; i < 3.05; i += 0.38 {
		result = laba4.Colcul(a, b, i)
		fmt.Printf("при x = %.2f, y = %.2f \n", i, result)
	}

	fmt.Println("\nзадача B:")
	for i := 0; i < len(x); i++ {
		result = laba4.Colcul(a, b, x[i])
		fmt.Println("при x =", x[i], " y =", result)
	}
}
