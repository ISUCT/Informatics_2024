package main

import (
	"fmt"

	"isuct.ru/informatics2022/laba4"
)

func main() {
	fmt.Println("Пучков Дмитрий Максимович")
	a := 2.25
	fmt.Println("\nзадача A:")
	xStartA := 1.2
	xEndA := 2.7
	step := 0.3
	for x := xStartA; x <= xEndA; x += step {
		y := laba4.Calcul(a, x)
		fmt.Printf(" x = %.2f, y = %.4f", x, y)
	}

	fmt.Println("\nзадача B:")
	xB := []float64{1.31, 1.39, 1.44, 1.56, 1.92}
	for _, x := range xB {
		y := laba4.Calcul(a, x)
		fmt.Printf(" x = %.2f, y = %.4f", x, y)
	}
}
