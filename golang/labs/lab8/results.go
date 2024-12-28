package lab8

import (
	"fmt"
	"strconv"
	"strings"

	"isuct.ru/informatics2022/labs/lab4"
)

const PathTask1 = "labs/lab8/input.txt"

func Task1() {
	text, err := ReadFile(PathTask1)
	if err != nil {
		panic(err)
	}

	var result []float64
	Parameters := strings.Split(text, "\n")
	for _, P := range Parameters {
		number, err := strconv.ParseFloat(P, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}

	b := result[0]
	begin_x := result[1]
	end_x := result[2]
	delta_x := result[3]

	fmt.Println("Значение A:")
	resultA := lab4.TaskA(b, begin_x, end_x, delta_x)

	for _, y := range resultA {
		fmt.Printf("y: %.4f\n", y)
	}

	fmt.Println("\nЗначение B:")
	resultB := lab4.TaskB(b, result[4:])

	for _, y := range resultB {
		fmt.Printf("y: %.4f\n", y)
	}
}
