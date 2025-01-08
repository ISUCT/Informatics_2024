package lab4

import (
	"fmt"
	"math"
	"strconv"

	"isuct.ru/informatics2022/lab8"
)

func Calculate(a,b,x float64) float64 {
	y := ((math.Log10(math.Pow(x,2)-1))/(math.Log(a*math.Pow(x,2)-b)/math.Log(5)))
	return y
}

func TaskA(a, b, x1, x2, dx float64) []float64 {
	var y []float64
	for x := x1; x<=x2; x += dx {
		y = append(y, Calculate(a,b,x))
	}
	return y
}

func TaskB(a float64, b float64, x[] float64) []float64 {
	var arr []float64
	for _, value := range x {
		arr = append(arr, Calculate(a,b, value))
	}
	return arr
}

func PrintValue(Values []float64) {
	for _, value := range Values {
		fmt.Println(value)
	}
}

func ReadFileTask4() []float64 {
	fmt.Println("Введите X1, X2, dX для задачи A и список значений для задачи B")
	arr := lab8.RunLab8()
	var numbers []float64
	for _, values := range arr {
		value, _ := strconv.ParseFloat(values, 64)
		numbers = append(numbers, value)
	}
	return numbers
}


func RunLab4 () {
	const a = 1.1
	const b = 0.09
	arr := ReadFileTask4()
	slice := arr[3:]

	ValuesA := TaskA(a, b, arr[0], arr[1], arr[2])
	ValuesB := TaskB(a, b, slice)

	PrintValue(ValuesA)
	PrintValue(ValuesB)
}
