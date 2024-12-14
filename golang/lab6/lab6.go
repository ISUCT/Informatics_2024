package lab6

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func (c *Person) SetAge(age int) {
	c.Age = age
}

func (c *Person) SetCountry(country string) {
	c.Country = country
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

	Yan.SetAge(31)

	Yan.SetCountry("Беларусь")

	Yan.DisplayInfo()
}
