package LAB_6

import (
	"fmt"
	"strconv"
)

const maxRabbitAge = 12

type Rabbit struct {
	Name  string
	Age   int
	Color string
}

func (rabbit *Rabbit) SetName(name string) {
	rabbit.Name = name
}

func (rabbit *Rabbit) SetColor(color string) {
	rabbit.Color = color
}
func (rabbit *Rabbit) SetAge(age int) {
	if age <= maxRabbitAge {
		rabbit.Age = age
	} else {
		fmt.Println("Ошибка: указанный возраст превышает максимальный возраст кроликов!")
	}
}

func (rabbit Rabbit) GetInfo() string {
	return "Имя кролика: " + rabbit.Name + "\nВозраст кролика: " + strconv.Itoa(rabbit.Age) + "\nЦвет кролика: " + rabbit.Color
}

func Lab6() {
	rabbit := Rabbit{Name: "Беляш", Age: 1, Color: "Белый"}

	fmt.Println(rabbit.GetInfo())

	rabbit.SetName("Пушок")
	rabbit.SetAge(3)

	fmt.Println(rabbit.GetInfo())
}
