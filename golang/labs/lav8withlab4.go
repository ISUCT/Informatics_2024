package labs

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func logarifmLab8(x, base float64) float64 {
	return math.Log(x) / math.Log(base)
}

func calculateYLab8(a, b, x float64) float64 {
	koren := math.Cbrt(x)
	log5 := logarifmLab8(x, 5)
	logValue := math.Log(x - 1)

	lgkub := math.Pow(logValue, 3)
	return (a*koren - b*log5) / lgkub
}

func taskBLab8(a, b float64, xValues []float64) []float64 {
	var yResults []float64
	for _, x := range xValues {
		y := calculateYLab8(a, b, x)
		yResults = append(yResults, y)
	}
	return yResults
}

func readInput(filename string) (float64, float64, []float64, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []float64

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return 0, 0, nil, err
		}
		numbers = append(numbers, num)
	}

	a := numbers[0]
	b := numbers[1]
	xValues := numbers[2:]

	return a, b, xValues, nil
}

func Runlab8with4() {
	a, b, xValues, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Ошибка в прочтении файла", err)
		return
	}

	yResults := taskBLab8(a, b, xValues)
	fmt.Println("X :", xValues)
	fmt.Println("Y :", yResults)
}
