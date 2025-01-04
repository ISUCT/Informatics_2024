package lab6

import "fmt"

type Person struct {
	Name   string
	Age    int
	Height float64
}

func NewPerson(name string, age int, height float64) *Person {
	p := new(Person)
	p.Name = name
	p.Age = age
	p.Height = height
	return p
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p Person) GetAge() int {
	return p.Age
}

func (p Person) GetHeight() float64 {
	return p.Height
}

func CompleteLab6() {
	ivan := NewPerson("Дима", 20, 1.78)
	ivan.SetAge(20)

	fmt.Printf("Имя: %s\nВозраст: %d\nРост: %.2f\n", ivan.Name, ivan.GetAge(), ivan.GetHeight())
}
