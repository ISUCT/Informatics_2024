package lab4

import (
	"fmt"
	"math"
	"strconv"

	"isuct.ru/informatics2022/lab/lab8"
)

func Calculate(b, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3)))
}

func TaskA(b, Xn, Xk, delX float64) []float64 {
	var arr []float64
	for x := Xn; x <= Xk; x += delX {
		arr = append(arr, Calculate(b, x))
	}
	return arr
}

func TaskB(b float64, x []float64) []float64 {
	var arr []float64
	for _, value := range x {
		arr = append(arr, Calculate(b, value))
	}
	return arr
}

func PrintValue(Values []float64) {
	for _, value := range Values {
		fmt.Println(value)
	}
}

func ReadFileTask4() []float64 {
	fmt.Println("Вводите данные в следующем порядке: Xn, Xk, delX для задачи A и slice для задачи B")
	var arr []string
	arr = lab8.RunLab8Tasks()
	var numbers []float64
	for _, values := range arr {
		value, _ := strconv.ParseFloat(values, 64)
		numbers = append(numbers, value)
	}
	return numbers
}

func RunLab4Tasks() {
	const b float64 = 2.5
	var arr []float64
	arr = ReadFileTask4()
	slice := arr[3:]
	ValuesA := TaskA(b, arr[0], arr[1], arr[2])
	ValuesB := TaskB(b, slice)
	PrintValue(ValuesA)
	PrintValue(ValuesB)
}
