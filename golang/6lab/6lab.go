package sixlab

import "fmt"

type Dog struct {
	Name  string
	Age   int
	Breed string
}

func NewDog(name string, age int, breed string) Dog {
	return Dog{
		Name:  name,
		Age:   age,
		Breed: breed,
	}
}
func (d *Dog) GetAge() int {
	return d.Age
}
func (d *Dog) SetAge(age int) {
	d.Age = age
}
func (d *Dog) Bark() string {
	return "Woof!"
}

func Start6lab() {
	dog := NewDog("Ster", 5, "KaNTeR_V")
	fmt.Println("Возраст собаки:", dog.GetAge())
	dog.SetAge(6)
	fmt.Println("Обновленный возраст собаки:", dog.GetAge())
	fmt.Println("Собака лает:", dog.Bark())
}
