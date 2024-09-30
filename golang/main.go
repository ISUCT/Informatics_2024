package main

import (
    "fmt"
    "math"
)

func calculateFunction(a, x float64) (float64, error) {
	if x * x <= 1 {
		return 0, fmt.Errorf("Некорректное значение x")
	}

	result1 := math.Pow(a, x * x - 1)
	result2 := math.Log10(x * x - 1)
	result3 := math.Cbrt(x * x - 1)

	result := result1 - result2 + result3

	return result, nil
}

func main() {
	
	a := 1.6
	x := 2.0
	
	result, err := calculateFunction(a, x)
	if err != nil {
		fmt.Println("Ошибка", err)
	} else {
		fmt.Println("Результат:", result)
	}
	
}


