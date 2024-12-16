package laba8

import (
	"io"
	"os"
	"strconv"
	"strings"

	"isuct.ru/informatics2022/laba4"
)

const PathTask1 = "laba8/input.txt"

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

	begin_x := result[0]
	end_x := result[1]
	delta_x := result[2]
	laba4.Y_A(begin_x, end_x, delta_x)
	laba4.Y_B(result[3:7])

	return nil
}
