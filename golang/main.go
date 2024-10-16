package main

import (
		"fmt"

		labs "isuct.ru/informatics2022/labs/lab4"
)
func main() {
	fmt.Println("Кузьминов Илья Владимирович")
	results := labs.Laba4A()
	for _, result := range results {
		fmt.Println(result)
	}
	labs.Laba4B()
}


