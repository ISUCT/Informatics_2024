package main

import (
	"fmt"

	"isuct.ru/informatics2022/Lab4"
)

func main() {
	fmt.Println("A")

	a := 0.8
	b := 0.4

	for x := 1.23; x <= 7.23; x += 1.2 {
		fmt.Println("y=", Lab4.Calculate(a, b, x))
	}

	fmt.Println("B")

	var xValue = []float64{1.88, 2.26, 3.84, 4.55, -6.21}

	for _, x := range xValue {
		fmt.Println("y=", Lab4.Calculate(a, b, x))
	}

}
