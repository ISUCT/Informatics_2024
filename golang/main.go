package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Print("Задача А\n")
	var x_a float64 = 1.28
	for i := 0; i < 6   ; i++ { 
		fmt.Print(lab4.Function(x_a), "\n")
		x_a = x_a + 0.4
	}

	fmt.Print("Задача В\n")
	xs := [5]float64{1.1, 2.4, 3.6, 1.7, 3.9} 
	for _, x_b := range xs {
		fmt.Print(lab4.Function(x_b), "\n")
  }
}
