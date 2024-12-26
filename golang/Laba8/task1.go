package lab8

import (
	"io"
	"os"
	"strconv"
	"strings"

	lab4 "isuct.ru/informatics2022/Laba4"
)

const PathTask1 = "Laba8/input.txt"

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
	Parameters := strings.Split(text, "\r\n")
	for _, P := range Parameters {
		number, err := strconv.ParseFloat(P, 64)
		if err != nil {
			return err
		}
		result = append(result, number)
	}

	a := result[0]
	b := result[1]
	xMin := result[2]
	xMax := result[3]
	xDelta := result[4]
	lab4.CompleteTaskA(a, b, xMin, xMax, xDelta)
	lab4.CompleteTaskB(a, b, result[6:])

	return nil
}
