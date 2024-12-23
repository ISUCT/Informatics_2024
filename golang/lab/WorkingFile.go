package lab

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func CreateFile(path string) error {
	file, errCreate := os.Create(path)

	if errCreate != nil {
		return errCreate
	}
	defer file.Close()
	return nil
}

func WriteFile(path string, constants []byte) error {
	file, errCreate := os.OpenFile(path, os.O_RDWR, 0666)

	if errCreate != nil {
		return errCreate
	}

	file.Write(constants)
	defer file.Close()
	return nil
}

func FindValue(path string) []float64 {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := make([]byte, 64)
	constant := ""

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		constant += string(data[:n])
	}

	example := `\d\.\d+`
	re := regexp.MustCompile(example)
	values := re.FindAllString(constant, -1)
	list := make([]float64, 0, len(values))

	for _, value := range values {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, f)
	}

	defer file.Close()
	return list
}
