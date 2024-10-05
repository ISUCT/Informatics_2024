package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	var exb = []float64{1.24, 1.38, 2.38, 3.21, 0.68}
	fmt.Println("Задача А")
	for x := 0.33; x <= 1.23; x += 0.18 {
		y := lab4.F(x)
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
	fmt.Println()
	fmt.Println("Задача В")
	for _, x := range exb {
		y := lab4.F(x)
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
}
