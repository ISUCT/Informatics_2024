package main

import (
	"fmt"

	"isuct.ru/informatics2022/labs/lab4"
	"isuct.ru/informatics2022/labs/lab6"
	"isuct.ru/informatics2022/labs/lab7"
	"isuct.ru/informatics2022/labs/lab8"
)

func main() {
	fmt.Println("Пшеничников Иван Михайлович")
	lab4.Runlab4()
	lab6.Runlab6()
	lab7.RunLab7()
	lab8.RunLab8WithLab4()
	lab8.RunLab8()
}
