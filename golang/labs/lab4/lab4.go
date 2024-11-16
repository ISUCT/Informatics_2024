package lab4

import (
	"fmt"
	"math"

	"isuct.ru/informatics2022/labs/lab8"
)

func Calculate(a, b, x float64) float64 {
	return (math.Cbrt(a) + math.Pow(math.Tan(b*x), 4.5)) / (math.Pow(b, 0.2) + (1 / (math.Pow(math.Tan(a*x), 2.7))))
}

func TaskA(a, b, xn, xk, deltax float64) [][]float64 {
	var taskAanswer [][]float64
	for x := xn; x < xk; x += deltax {
		taskAanswer = append(taskAanswer, []float64{x, Calculate(a, b, x)})
	}
	return taskAanswer
}

func TaskB(a, b float64, taskBslice []float64) [][]float64 {
	var taskBanswer [][]float64
	for _, x := range taskBslice {
		taskBanswer = append(taskBanswer, []float64{x, Calculate(a, b, x)})
	}
	return taskBanswer
}

func RunLab4() {
	values, err := lab8.ReadDataFromFile("input.txt")
	if err != nil {
		panic("ошибка при чтении данных из файла")
	}
	var a float64 = values[0]
	var b float64 = values[1]
	var taskBslice = values[5:]
	fmt.Println("Задача А")
	var xn float64 = values[2]
	var xk float64 = values[3]
	var deltax float64 = values[4]
	for _, pair := range TaskA(a, b, xn, xk, deltax) {
		x := pair[0]
		y := pair[1]
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
	fmt.Println()
	fmt.Println("Задание B")
	for _, pair := range TaskB(a, b, taskBslice) {
		x := pair[0]
		y := pair[1]
		fmt.Printf("x = %.2f\ty = %f\n", x, y)
	}
}
