package main

import "isuct.ru/informatics2022/lab4"

func main() {
	lab4.TaskA(0.22, 0.92, 0.14)                  //Значения из таблицы
	var s = [5]float64{0.1, 0.35, 0.4, 0.55, 0.6} //Тут тоже значения из таблицы
	lab4.TaskB(s)
}
