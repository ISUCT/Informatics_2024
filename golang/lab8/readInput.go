package lab8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ParamsPath = "lab8/input.txt"

func ReadFileForLab4() ([]float64, error) {
	link := ParamsPath

	data, err := os.ReadFile(link)
	if err != nil {
		return nil, fmt.Errorf("(ReadFileForLab4) ошибка при чтении файла %s: %w", link, err)
	}

	var result []float64
	listParameters := strings.Split(string(data), "\n")

	for _, l := range listParameters {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		i, errParseFloat := strconv.ParseFloat(l, 64)
		if errParseFloat != nil {
			return nil, fmt.Errorf("(ReadFileForLab4) ошибка преобразования строки в число: %w", errParseFloat)
		}
		result = append(result, i)
	}

	return result, nil
}
