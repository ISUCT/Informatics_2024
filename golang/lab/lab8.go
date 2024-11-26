package lab

import (
	"fmt"
	"os"
)

func Lab8TasksA() {
	var number string
	file, err := os.Create("lab4A.txt")
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 3; i++ {
		fmt.Print("Введите ", i, " значение: ")
		fmt.Fscan(os.Stdin, &number)
		file.WriteString(number + ",")
	}
	defer file.Close()
}

func Lab8TasksB() {
	var number string
	file, err := os.Create("lab4B.txt")
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 5; i++ {
		fmt.Print("Введите ", i, " значение: ")
		fmt.Fscan(os.Stdin, &number)
		file.WriteString(number + ",")
	}
	defer file.Close()
}

func RunLab8Tasks() {
	fmt.Println("Введите значения Xn, Xk, delX для задачи A")
	Lab8TasksA()
	fmt.Println("Введите 5 числовых значений для задачи B")
	Lab8TasksB()
}
