package lab6

import (
	"fmt"
)

type Rabbit struct {
	Name   string
	Age    int
	Weight float64
}

func (rabbit *Rabbit) SetAge(age int) {
	if age > 0 {
		rabbit.Age = age
	}
}

func (rabbit *Rabbit) SetName(name string) {
	rabbit.Name = name
}

func (rabbit *Rabbit) SetWeight(weight float64) {
	if weight > 0 {
		rabbit.Weight = weight
	}
}

func (rabbit Rabbit) GetInfo() string {
	return fmt.Sprintf("Имя Кролика: %s, Возраст Кролика: %d, Вес Кролика: %.2f кг", rabbit.Name, rabbit.Age, rabbit.Weight)
}

func RunLab6() {
	var rabbit = Rabbit{Name: "Степан", Age: 1, Weight: 5}

	fmt.Println(rabbit.GetInfo())

	rabbit.SetAge(3)

	fmt.Println(rabbit.GetInfo())
}
