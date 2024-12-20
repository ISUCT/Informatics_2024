package lab8

import (
	"fmt"
	"strconv"
	"strings"
)

const LincTask1 = "Lab8/input.txt"

func ChangeStringToNumber(text string) ([]float64, error) {
	var list []float64
	Tasks := strings.Split(text, "\r\n")
	fmt.Println(Tasks)
	for _, v := range Tasks {
		number, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		list = append(list, number)
	}
	return list, nil
}
