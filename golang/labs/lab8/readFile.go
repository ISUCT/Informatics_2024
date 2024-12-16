package lab8

import (
	"fmt"
	"io"
	"os"
)

func read(str string) {
	file, err := os.Open(str)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
}
