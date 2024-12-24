package Lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	A        float64
	B        float64
	XValues  []float64
	ResultsA []float64
	ResultsB []float64
}

func calcY(b, x float64) float64 {
	denominator := math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3))
	if denominator == 0 {
		return math.Inf(1)
	}
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / denominator
}

func taskA(a, xStart, xEnd, deltaX float64) []float64 {
	var results []float64
	for x := xStart; x <= xEnd; x += deltaX {
		y := calcY(a, x)
		results = append(results, y)
	}
	return results
}

func taskB(b float64, xValues []float64) []float64 {
	var results []float64
	for _, x := range xValues {
		y := calcY(b, x)
		results = append(results, y)
	}
	return results
}

func RunLab8() {
	const inputFilename = "input.txt"
	const outputFilename = "output.txt"

	values, err := readInputFile(inputFilename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	if len(values) < 2 {
		fmt.Println("Недостаточно значений в файле")
		return
	}

	task := Task{
		A:       values[0],
		B:       values[1],
		XValues: values[2:],
	}

	xStart := 1.28
	xEnd := 3.28
	deltaX := 0.4

	task.ResultsA = taskA(task.A, xStart, xEnd, deltaX)
	fmt.Println("Результаты задания A:")
	for i, y := range task.ResultsA {
		x := xStart + float64(i)*deltaX
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}

	task.ResultsB = taskB(task.B, task.XValues)
	fmt.Println("\nРезультаты задания B:")
	for i, y := range task.ResultsB {
		x := task.XValues[i]
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}

	err = writeOutputFile(outputFilename, task.ResultsA, task.ResultsB, xStart, deltaX, task.XValues)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	err = displayFileContents(outputFilename)
	if err != nil {
		fmt.Println("Ошибка при выводе данных из файла:", err)
		return
	}

	fmt.Print("Введите текст для поиска в файле: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchTerm := scanner.Text()
	err = searchInFile(outputFilename, searchTerm)
	if err != nil {
		fmt.Println("Ошибка при поиске в файле:", err)
		return
	}
}

func readInputFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var values []float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Ошибка при чтении числа:", err)
			continue
		}
		values = append(values, value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return values, nil
}

func writeOutputFile(filename string, resultsA, resultsB []float64, xStart, deltaX float64, xValues []float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("Результаты задания A:\n")
	for i, y := range resultsA {
		x := xStart + float64(i)*deltaX
		writer.WriteString(fmt.Sprintf("x: %.2f, y: %.4f\n", x, y))
	}

	writer.WriteString("\nРезультаты задания B:\n")
	for i, y := range resultsB {
		x := xValues[i]
		writer.WriteString(fmt.Sprintf("x: %.2f, y: %.4f\n", x, y))
	}

	writer.Flush()
	return nil
}

func displayFileContents(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("\nСодержимое файла:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func searchInFile(filename, searchTerm string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Printf("Результаты поиска для '%s':\n", searchTerm)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchTerm) {
			fmt.Println(line)
			found = true
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if !found {
		fmt.Println("Совпадений не найдено.")
	}
	return nil
}
