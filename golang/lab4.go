package main

import (
	"fmt"
	"math"
)

const a = 1.2
const b = 0.48


func main() {
    fmt.Printf("task1: %v\n", task1())
    fmt.Printf("task2: %v\n", task2())
}

func task1 () []float64 {
	const xN, xK, dX = 0.7, 2.2, 0.3
	var result []float64
	for i:= xN; i <= xK; i = i + dX {
		result = append(result, computeYFromX(a, b, i))
	}
	return result
}

func task2 ()[]float64 {
	var xInputs = [5]float64 {0.25, 0.35, 0.56, 0.94, 1.28}
	var result []float64
	for _, element := range xInputs {
		result = append(result, computeYFromX(a, b, element))
	}
	return result
}

func computeYFromX (a float64, b float64, x float64) float64 {
    var bCube = math.Pow(b, 3)
    var sinAXSquare = math.Pow(math.Sin(a * x), 2)
    var acoxXBX = math.Acos(x*b*x)
    var powE = math.Pow(math.E, (-x/2))
    return (bCube + sinAXSquare )/(acoxXBX + powE)
}