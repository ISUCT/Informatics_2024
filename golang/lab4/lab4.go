
package lab4 
 
import ( 
    "math"
    "fmt"
) 
 
func  CalculateExpression (a float64, x float64) float64 { 
    return math.Tan(math.Pow(math.Log10(a+x), 3)) / math.Pow(a+x, 2.0/7.0) 
} 
 
func TaskA(a , xi , xk , deltax float64) []float64 { 
     var values []float64 
    for x := xi; x <= xk; x += deltax { 
        values = append(values,  CalculateExpression (a, x)) 
    } 
    return values 
} 
 
func TaskB(a float64, x [5]float64) []float64 { 
    var values []float64
    for _, value := range x {
        values = append(values,  CalculateExpression (a, value)) 
    } 
    return values
} 


func RunLab4Tasks() {
    a := 2.0
    fmt.Println(TaskA(a, 1.08, 1.88, 0.16))
	var s = [5]float64{1.16, 1.35, 1.48, 1.52, 1.96}
	fmt.Println(TaskB(a,s))
}
