package main

import (
	"isuct.ru/informatics2022/lab4"
	"fmt"
	"math"
)



func main() {
	a := -2.5

	//Задача A
	fmt.Println("Задача A:")
xН := 3.5
xК := 6.5
deltax := 0.6

for x := xН; x <= xК; x += deltax {
	y := lab4.CalculateY(x, a)
	fmt.Printf("x = %.2f, y = %.4f\n", x, y)
}

fmt.Println("Задача B:")
xValues := []float64{2.89, 3.54, 5.21, 6.28, 3.48}

for _, x:= range xValues {
y := lab4.CalculateY(x, a)
fmt.Printf("x = %.2f, y = %.4f\n", x, y)
}
}
