package lab8

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CompleteLaba8() {
	path := "lab8/lab8.txt"
	CreateFile(path)
	WriteFile(path)
	ReadFile(path)
}

func CreateFile(path string) {
	_, errStat := os.Stat(path)
	if errStat != nil {

		file, errCreate := os.Create(path)
		if errCreate != nil {
			log.Fatal(errCreate)
			os.Exit(1)
		}
		file.Close()
	}
}

func WriteFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Print("Введите текст: ")
	text, _ := in.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	file.WriteString(text + " ")
}

func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var result string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		result = string(data[:n])
	}
	fmt.Print(result)
}
