package lab8

import (
	"io"
	"os"
	"strconv"
	"strings"

	"isuct.ru/informatics2022/laba4"
)

const PathTask1 = "lab8/input.txt"

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

	a := result[0]
	b := result[1]
	xMin := result[2]
	xMax := result[3]
	xDelta := result[4]
	laba4.TaskA(a, b, xMin, xMax, xDelta)
	laba4.TaskB(a, b, result[5:9])

	return nil
}
