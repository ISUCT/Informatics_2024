package lab6

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func (c *Person) SetAge() {
	c.Age = 31
	fmt.Printf("%s теперь твой возраст %d!\n", c.Name, c.Age)
}

func (c *Person) DisplayInfo() {
	fmt.Println("Имя:", c.Name)
	fmt.Println("Возраст:", c.Age)
	fmt.Println("Страна:", c.Country)
}

func Lab6() {
	Yan := Person{
		Name:    "Ян",
		Age:     20,
		Country: "Россия",
	}

	Yan.DisplayInfo()

	Yan.SetAge()

	Yan.DisplayInfo()
}
