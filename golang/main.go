package main

import (
	"fmt"

	"isuct.ru/informatics2022/labs/lab4"

	"isuct.ru/informatics2022/labs/lab6"

	"isuct.ru/informatics2022/labs/lab7"

	"isuct.ru/informatics2022/labs/lab8"
)

func main() {
	fmt.Println("Васильев Никита Глебович")
	fmt.Println("4 лабораторная:")
	lab4.RunLab4()
	fmt.Println("6 лабораторная:")
	lab6.RunLab6()
	fmt.Println("7 лабораторная:")
	lab7.RunLab7()
	fmt.Println("Исправленная 4 лаба (1 часть 8 лабы)")
	lab8.RunFixedLab4()
	fmt.Println("8 лабораторная, часть 2")
	lab8.RunLab8()
}
