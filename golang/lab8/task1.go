package lab8

import (
	"fmt"
	"os"
)

const Numbers = "lab8/input.txt"

func ReadNumbers() {
	data, err := os.Open("lab8/input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	dataContent, err := os.ReadFile("lab8/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataContent))
}
