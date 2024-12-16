package lab8

import (
	"fmt"

	"io"

	"os"

	"strconv"

    "isuct.ru/informatics2022/labs/lab4"
)
 
func getNumbers() (b, xbegin, xend, xdelta float64, xs []float64) {
    file, err := os.OpenFile("labs/lab8/numbersFromLab4", os.O_RDONLY, 0666)
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 

    data := make([]byte, 1)
    var buf []string
    var str string
    for{
        
        _, err := file.Read(data)
        if err == io.EOF{  
            buf = append(buf, str)
            str = ""
            break
        } else if string(data) == "\n" {
            buf = append(buf, str)
            str = ""
            continue
        }
        str += string(data)
    }

    b, _  = strconv.ParseFloat(buf[0], 64)
    buf = append(buf[:0], buf[1:]...)
    xbegin, _  = strconv.ParseFloat(buf[0], 64)
    buf = append(buf[:0], buf[1:]...)
    xend, _  = strconv.ParseFloat(buf[0], 64)
    buf = append(buf[:0], buf[1:]...)
    xdelta, _  = strconv.ParseFloat(buf[0], 64)
    buf = append(buf[:0], buf[1:]...)
    for _, num := range buf {
        numfloat, _ := strconv.ParseFloat(num, 64)
        xs = append(xs, numfloat)
    }
    return 
}

func ansTask1 () {
    b, xbegin, xend, xdelta, xs := getNumbers()
	var for_a = lab4.CompliteTaskA(b, xbegin, xend, xdelta)
	var for_b = lab4.CompliteTaskB(b, xs)

	fmt.Println("Решение задания А:")
	for i, y := range for_a {
		fmt.Printf("y%d = %f\n", i+1, y)
	}

	fmt.Println("Решение задания B:")
	for i, y := range for_b {
		fmt.Printf("y%d = %f\n", i+1, y)
	}
}