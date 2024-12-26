package lab8

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	lab4 "isuct.ru/informatics2022/Lab_4"
)

const PathTask1 = "Lab_8/input.txt"

func task1() error {
	file, err := os.Open(PathTask1)
	if err != nil {
		return err
	}
	defer file.Close()

	var text string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		text = string(data[:n])
	}

	var result []float64
	Parameters := strings.Split(text, "\n")
	for _, P := range Parameters {
		number, err := strconv.ParseFloat(P, 64)
		if err != nil {
			return err
		}
		result = append(result, number)
	}

	a := 2.0
	b := 1.5
	begin_x := result[0]
	end_x := result[1]
	delta_x := result[2]
	fmt.Println(lab4.TaskA(a, b, begin_x, end_x, delta_x))
	fmt.Println(lab4.TaskB(a, b, result[3:7]))

	return nil
}
