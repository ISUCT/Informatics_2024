package lab4

import ( 
	"fmt"
	"math"
)

func ColculateFunction(a, b, x float64) float64 {
  if x == 0 {
    return 0
  }
 var y = math.Sqrt(float64(math.Abs(a-b*x)) / math.Pow(math.Log10(float64(x)), 3))

 return y;
}

func completeTaskA(a, b, xn, xk, xd float64) []float64 {
	var result []float64
	for i := xn; i < xk; i += xd {
	result = append(result, ColculateFunction(a, b, i))
	}
	return result
}
func completeTaskB(a, b float64, x []float64) []float64 {
	var result []float64
	for _, i := range x {
	result = append(result, ColculateFunction(a, b, i))
	}
	return result
}
func CompleteLab4() {
	var a float64 = 7.2
	var b float64 = 4.2
	var x []float64 = []float64{2.4, 2.8, 3.9, 4.7, 3.16}
	var xn float64 = 1.81
	var xk float64 = 5.31
	var xd float64 = 0.7
	var resultA []float64 = completeTaskA(a, b, xn, xk, xd)
	fmt.Println(resultA)
	var resultB []float64 = completeTaskB(a, b, x)
	fmt.Println(resultB)
}