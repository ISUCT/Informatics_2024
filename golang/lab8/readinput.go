package lab8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ParamsPath = "lab8/input.txt"

func ReadFileForLab4() ([]float64, error) {
	var link string = ParamsPath
	data, err := os.ReadFile(link)
	if err != nil {
		return nil, fmt.Errorf("(ReadFileForLab4) чтение файла %s: %w", link, err)
	}

	var result []float64
	listParameters := strings.Split(string(data), "\r\n")
	for _, l := range listParameters {
		i, errParseFloat := strconv.ParseFloat(l, 64)
		if errParseFloat != nil {
			return nil, fmt.Errorf("(ReadFileForLab4) строка->число: %w", errParseFloat)
		}
		result = append(result, i)
	}

	return result, nil
}
