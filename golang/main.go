package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	var res float64
	var elements = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	fmt.Println("Задача A \n")
	for i := 0; i < len(elements); i++ {
		res = lab4.Mathfunc(7.2, 1.3, elements[i])
		fmt.Println("При x =", elements[i], "функция равна: \n", res)
	}
	fmt.Println("Задача B \n")
	for i := 1.56; i <= 4.71; i += 0.63 {
		res = lab4.Mathfunc(7.2, 1.3, i)
		fmt.Println("При x =", i, "функция равна: \n", res)
	}
	fmt.Println("Нерабеев Кирилл Сергеевич")
}
