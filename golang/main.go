package main

import (
	"fmt"

	lab4 "isuct.ru/informatics2022/labs"
)

func main() {
	var exb = []float64{1.24, 1.38, 2.38, 3.21, 0.68}
	fmt.Println("Задача А")
	for x := 0.33; x <= 1.23; x += 0.18 {
		y := lab4.F(x)
		fmt.Println(x, y)
		fmt.Printf("x = %f\ty = %f\n", x, y)
	}
	fmt.Println()
	fmt.Println("Задача В")
	for _, x := range exb {
		y := lab4.F(x)
		fmt.Println(x, y)
		fmt.Printf("x = %f\ty = %f\n", x, y)
	}
}
