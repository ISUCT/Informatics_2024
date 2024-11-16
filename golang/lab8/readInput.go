package lab8

import (
	"os"
	"strconv"
	"strings"
)

func ReadFileForLab4() ([]float64, error) {
	var link string = "lab8/input.txt"
	data, err := os.ReadFile(link)
	if err != nil {
		return nil, err
	}

	var result []float64
	line := strings.Split(string(data), "\r\n")
	for _, l := range line {
		i, _ := strconv.ParseFloat(l, 64)
		result = append(result, i)
	}

	return result, nil
}
