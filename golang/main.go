package main

import (
	"fmt"

	Lab4 "isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Print("Задача А\n")
	var xn, xk, xdel float64 = 0.26, 0.66, 0.08
	for x := xn; x <= xk; x += xdel {
		fmt.Println(Lab4.Сalcul(x))
	}

	fmt.Print("Задача B\n")
	var xv = [5]float64{0.1, 0.35, 0.4, 0.55, 0.6}
	for _, x := range xv {
		fmt.Println(Lab4.Сalcul(x))
	}
}
