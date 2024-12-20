package lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func logarifm(x, base float64) float64 {
	return math.Log(x) / math.Log(base)
}

func calculateY(a, b, x float64) float64 {
	koren := math.Cbrt(x)
	log5 := logarifm(x, 5)
	logValue := math.Log(x - 1)

	if logValue == 0 {
		return math.NaN()
	}

	lgkub := math.Pow(logValue, 3)
	return (a*koren - b*log5) / lgkub
}

func taskA(a, b, xn, xk, xd float64) ([]float64, []float64) {
	var xResults, yResults []float64
	for x := xn; x <= xk; x += xd {
		y := calculateY(a, b, x)
		xResults = append(xResults, x)
		yResults = append(yResults, y)
	}
	return xResults, yResults
}

func taskB(a, b float64, xValues []float64) []float64 {
	var yResults []float64
	for _, x := range xValues {
		y := calculateY(a, b, x)
		yResults = append(yResults, y)
	}
	return yResults
}

func ReadTxt(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла %s: %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []float64

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("Ошибка при конвертации'%s': %w", line, err)
		}
		values = append(values, value)
	}

	if len(values) != 10 {
		return nil, fmt.Errorf("Недостаточно данных, нужно 10 значений")
	}

	return values, nil
}

func RunLab8WithLab4() {
	values, err := ReadTxt("labs/lab8/input.txt")
	if err != nil {
		panic("Ошибка при чтении данных")
	}

	a := values[0]
	b := values[1]
	xn := values[2]
	xk := values[3]
	dx := values[4]
	arguments := values[5:]
	fmt.Print("A:")
	fmt.Println(taskA(a, b, xn, xk, dx))

	fmt.Print("B:")
	fmt.Println(taskB(a, b, arguments))
}
