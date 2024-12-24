package lab

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func ShowLab8() {
	path := "lab/input.txt"
	CreateFile(path)
	var data []byte

	input := InputValue()
	data = []byte(input)

	WriteFile(path, data)
	fmt.Println("Файл input.txt успешно создан")

	var A []float64
	var B []float64

	list := FindValue(path)

	a := list[0]
	b := list[1]

	A = lab4.TaskA(list[2], list[3], list[4])
	B = list[5:]

	fmt.Println("Задача A:")
	lab4.Enter(a, b, A)

	fmt.Println("Задача B:")
	lab4.Enter(a, b, B)
}
