package main

import (
	"fmt"

	"isuct.ru/informatics2022/labs/lab4"
	"isuct.ru/informatics2022/labs/lab6"
)

func main() {
	fmt.Printf("Бейлин Алексей Борисович")

	lab4.Ans_lab4()

	var cat = lab6.Cat{Name: "Имя", Age: 18, Breed: "Парода"}

	fmt.Println(lab6.Cat.GetStatus(cat))
	
	cat.SetName("Мурзик")
	cat.SetAge(10)
	cat.SetBreed("Сфинкс")

	fmt.Println(cat.GetStatus())
}
