package main

import (
	"fmt"

	"isuct.ru/informatics2022/function"
)

func main() {
	var b float64 = 2.5
	var xs = [8]float64{1.28, 3.28, 0.4, 1.1, 2.4, 3.6, 1.7, 3.9}
	for _, x := range xs {
		fmt.Println("y =", function.Uravnenie(b, x))
	}

}
