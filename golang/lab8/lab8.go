package lab8

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func MakeNewFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func InputData() []string {
	var count int
	fmt.Print("Введите количество значений: ")
	fmt.Fscan(os.Stdin, &count)
	var arr []string
	var value string
	for i := 1; i <= count; i++ {
		fmt.Print("Введите значение: ")
		fmt.Fscan(os.Stdin, &value)
		arr = append(arr, value)
	}
	return arr
}

func ReadInputFromFile(name string) []string {
	var arr []string
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for {
		data := make([]byte, 64)
		for {
			n, err := file.Read(data)
			if err == io.EOF {
				break
			}
			arr = append(arr, string(data[:n]))
			break
		}
		return arr
	}
}

func WriteInFile(name string, arr []string) {
	file, err := os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, value := range arr  {
		file.WriteString(value + "\n")
	}
}


func SearchInputFromFile(name string, ValueFile string) {
	arr := strings.Split( ReadInputFromFile(name)[0], "\n")
	for _, value := range arr {
		if value == ValueFile {
			fmt.Println("Ошибка при поиске в файле")
			break
		}
	}
}

func RunLab8Tasks() []string {
	var NameFile string
	fmt.Print("Введите название файла: ")
	fmt.Fscan(os.Stdin, &NameFile)
	MakeNewFile(NameFile)

	arr := InputData()
	WriteInFile(NameFile, arr)

	arr = strings.Split(ReadInputFromFile(NameFile)[0], "\n")
	arr = arr[0 : len(arr)-1]

	var Search string
	fmt.Print("Введите значение для поиска: ")
	fmt.Fscan(os.Stdin, &Search)
	SearchInputFromFile(NameFile, Search)

	return arr
}
