package lab6

import "fmt"

type Person struct {
	Name   string
	Age    string
	Height string
}

func NewPerson(name, age, height string) *Person {
	p := new(Person)
	p.Name = name
	p.Age = age
	p.Height = height
	return p
}

func (p *Person) SetName(name string) { p.Name = name }

func (p *Person) UpdateStruct(new_name, new_age, new_height string) {
	(*p).Name = new_name
	(*p).Age = new_age
	(*p).Height = new_height
}

func (p Person) GetName() string   { return p.Name }
func (p Person) GetAge() string    { return p.Age }
func (p Person) GetHeight() string { return p.Height }

func RunLab6Tasks() {
	human := NewPerson("Марк", "21", "182")
	human.SetName("Виталик")
	fmt.Println(human.GetName())
	fmt.Println(human.GetAge())
	fmt.Println(human.GetHeight())
	human.UpdateStruct("Мария", "19", "168")
	fmt.Println(human.GetName())
	fmt.Println(human.GetAge())
	fmt.Println(human.GetHeight())
}
