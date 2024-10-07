package main

import ("fmt"
"isuct.ru/informatics2022/lab4")

func main() {
	fmt.Println("Morozov_ilya_Dmitrievich")
	// Задаем начальное значение, конечное значение и шаг
	start := 0.26
	end := 0.66
	step := 0.4

	// Интерируем по значениям x
	for x := start; x<= end; x+= step {
		//Вычисляем arcsin и arccos
		y := lab4.Calculate(x) 
		

		// Выводим результат
		fmt.Println("x: %.2а, y: %.6f\n", x, y)
	}
}
