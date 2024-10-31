package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/labs/lab4"
	"isuct.ru/informatics2022/internal/labs/lab6"
)

func main() {
	fmt.Println("Semenenko Mikhail")
	resultA := lab4.Task_A(0.22, 0.92, 0.14)
	fmt.Println(resultA)
	resultB := lab4.Task_B(resultA)
	fmt.Println(resultB)

	lab6.RunLab6()
}
