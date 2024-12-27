package lab4

import (
    "fmt"
    "math"
    "strconv"
    "isuct.ru/informatics2022/lab8"
)

func CalculateExpression(a float64, x float64) float64 {
    return math.Tan(math.Pow(math.Log10(a+x), 3)) / math.Pow(a+x, 2.0/7.0)
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
    a := 2.0
    fmt.Println(TaskA(a, 1.08, 1.88, 0.16))
    
    var s = [5]float64{1.16, 1.35, 1.48, 1.52, 1.96}
    fmt.Println(TaskB(a, s))

    arr := GetInputForTask4()
    values := [5]float64{}
    
    copy(values[:], arr[3:8]) // Если arr имеет меньше элементов, это вызовет панику

    ValuesA := TaskA(a, arr[0], arr[1], arr[2])
    ValuesB := TaskB(a, values)
	
    fmt.Println(ValuesA)
    fmt.Println(ValuesB)
}
