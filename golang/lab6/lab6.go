package lab6

import "fmt"

type Dog struct {
	Name  string
	Age   int
	Breed string
}

func NewDog(name string, age int, breed string) *Dog {
	return &Dog{
		Name:  name,
		Age:   age,
		Breed: breed,
	}
}

func (d *Dog) Bark() string {
	return "Woof"
}

func (d *Dog) GetAge() int {
	return d.Age
}

func (d *Dog) SetAge(age int) {
	d.Age = age
}

func RunLab6() {
	dog := NewDog("Rex", 8, "Great Dane")
	fmt.Println(dog.Bark())
	fmt.Println(dog.GetAge())
	dog.SetAge(10)
	fmt.Println(dog.GetAge())
}
