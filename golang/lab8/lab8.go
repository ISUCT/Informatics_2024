package lab8

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)
func CreateNewFile(name string) {
    file, err := os.Create(name)
    if err != nil {
        panic(err)
    }
    defer file.Close()
}
func InputData() (float64, []float64) {
    var a float64
    fmt.Print("Введите значение a: ")
    fmt.Scan(&a)

    var n int
    fmt.Print("Введите количество значений x: ")
    fmt.Scan(&n)

    xValues := make([]float64, n)
    for i := 0; i < n; i++ {
        fmt.Printf("Введите значение x[%d]: ", i+1)
        fmt.Scan(&xValues[i])
    }

    return a, xValues
}
func WriteValuesToFile(filename string, a float64, xValues []float64) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(fmt.Sprintf("%f\n", a))
    if err != nil {
        return err
    }

    for _, x := range xValues {
        _, err = file.WriteString(fmt.Sprintf("%f\n", x))
        if err != nil {
            return err
        }
    }
    return nil
}

func ReadInputFromFile(filename string) (float64, []float64, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var a float64
    var xValues []float64

    if scanner.Scan() {
        a, err = strconv.ParseFloat(scanner.Text(), 64)
        if err != nil {
            return 0, nil, err
        }
    }

    for scanner.Scan() {
        x, err := strconv.ParseFloat(scanner.Text(), 64)
        if err != nil {
            return 0, nil, err
        }
        xValues = append(xValues, x)
    }

    return a, xValues, scanner.Err()
}
func PrintFileContent(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    return scanner.Err()
}

func SearchInFile(filename string, searchString string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var results []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if contains(line, searchString) {
            results = append(results, line)
        }
    }
    return results, scanner.Err()
}
func contains(line string, searchString string) bool {
    return strings.Contains(line, searchString)
}
  
