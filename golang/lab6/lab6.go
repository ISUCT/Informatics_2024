package lab6

import (
	"fmt"
	"strconv"
)

type Rabbit struct {
	Name  string
	Age   int
	Color string
}

func (rabbit *Rabbit) SetAge(age int) {
	rabbit.Age = age
}

func (rabbit *Rabbit) SetName(name string) {
	rabbit.Name = name
}

func (rabbit *Rabbit) SetColor(color string) {
	rabbit.Color = color
}

func (rabbit Rabbit) GetInfo() string {
	return_str := "Имя кролика: " + rabbit.Name + "\nВозраст кролика: " + strconv.Itoa(rabbit.Age) + "\nЦвет кролика: " + rabbit.Color
	return return_str
}

func RunLab6() {
	var rabbit = Rabbit{Name: "Имя", Age: 0, Color: "Цвет"}

	rabbit.SetName("Стеша")
	rabbit.SetAge(2)
	rabbit.SetColor("Рыжий")

	fmt.Println(rabbit.GetInfo())
}
