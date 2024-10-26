package main

import (
	"fmt"

	"isuct.ru/informatics2022/labs/lab4"
	"isuct.ru/informatics2022/labs/lab6"
)

func main() {
	fmt.Printf("Бейлин Алексей Борисович")

	lab4.AnsLab4()

	var cat = lab6.Cat{Name: "Имя", Age: 18, Breed: "Парода"}

	fmt.Println(cat.GetInfo())
	
	cat.SetName("Мурзик")
	cat.SetAge(10)
	cat.SetBreed("Сфинкс")

	fmt.Println(cat.GetInfo())
}
