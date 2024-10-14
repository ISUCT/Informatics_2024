package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	var res float64
	var elements = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	fmt.Println("Задача B \n")
	for _, i := range elements {
		res = lab4.Mathfunc(7.2, 1.3, i)
		fmt.Println("При x =", i, "функция равна: \n", res)
	}
	fmt.Println("Задача А \n")
	var xn, xk, xdel float64 = 1.56, 4.71, 0.63
	for i := xn; i <= xk; i += xdel {
		fmt.Println("При x =", i, "функция равна: \n", lab4.Mathfunc(7.2, 1.3, i))
	}
	fmt.Println("Нерабеев Кирилл Сергеевич")
}
