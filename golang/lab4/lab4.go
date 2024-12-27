package lab4 
 
import ( 
    "math"
    "fmt"
	"strconv"
    "isuct.ru/informatics2022/lab8"
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
func GetInputForTask4() []float64 {
	fmt.Println("Введите данные для Xi, Xk, delX и массив для задач A и B")
	arr := lab8.RunLab8Tasks()
	var input []float64
	for _, values := range arr {
		value, _ := strconv.ParseFloat(values, 64)
		input = append(input, value)
	}
	return input
}

func RunLab4Tasks() {
	const a float64 = 2.0
	arr :=GetInputForTask4()
	values := arr[3:]
	ValuesA := TaskA(a, arr[0], arr[1], arr[2])
	ValuesB := TaskB(a, values)
	PrintValue(ValuesA)
	PrintValue(ValuesB)
}
