package main

import (
	"fmt"

	"isuct.ru/informatics2022/laba4"
	"isuct.ru/informatics2022/laba6"
)

func main() {
	fmt.Println("Ткачёв Михаил Александрович")

	var a float64 = 2.5
	var b float64 = 4.6
	var xMin float64 = 1.15
	var xMax float64 = 3.05
	var xDelta float64 = 0.38
	var x []float64 = []float64{1.2, 1.36, 1.57, 1.93, 2.25}

	var resultA []float64 = laba4.TaskA(a, b, xMin, xMax, xDelta)
	fmt.Println(resultA)
	var resultB []float64 = laba4.TaskB(a, b, x)
	fmt.Println(resultB)

	//лабороторная работа 6
	chevrolet := laba6.NewCar("Шевроле", "синий", 60.0)
	chevrolet.SetSpeed(80.0)
	fmt.Println(chevrolet.GetSpeed())
	fmt.Println(chevrolet.GetColor())
}
