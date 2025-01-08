package lab6

import "fmt"

type Rabbit struct {
	Name  string
	Age   int
	Color string
}

func NewRabbit(name string, age int, color string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Name = name
	rabbit.Age = age
	rabbit.Color = color
	return rabbit
}

func (r Rabbit) GetName() string {
	return r.Name
}

func (r *Rabbit) SetAge(age int) {
	r.Age = age
}

func (r Rabbit) GetAge() int {
	return r.Age
}

func Lab6() {
	rabbit := NewRabbit("Шнафи", 3, "Чёрный")
	fmt.Println(rabbit.GetAge())
	fmt.Println(rabbit.GetName())
	rabbit.SetAge(4)
	fmt.Println(rabbit.GetAge())
}
