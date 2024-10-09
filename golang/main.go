package main

import (
	"isuct.ru/informatics2022/lab4"
	"fmt"
	"math"
)



func main() {
	//Параметры функции.
	a := -2.5

	//Задача A
	fmt.Println("Задача A:")
xН := 3.5  // начальное знакчение x
xК := 6.5  // конечное значение x
deltax := 0.6  // шаг изменения x

for x := xН; x <= xК; x += deltax {
	y := lab4.CalculateY(x, a)
	fmt.Printf("x = %.2f, y = %.4f\n", x, y)
}

//Задача В
fmt.Println("Задача B:")
xValues := []float64{2.89, 3.54, 5.21, 6.28, 3.48}

for _, x:= range xValues {
y := lab4.CalculateY(x, a)
fmt.Printf("x = %.2f, y = %.4f\n", x, y)
}
}
