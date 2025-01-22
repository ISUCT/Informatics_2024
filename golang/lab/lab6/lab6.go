package main

import "fmt"

type Dog struct {
	Name  string
	Breed string
	Age   int
}

func NewDog(name string, breed string, age int) Dog {
	return Dog{Name: name, Breed: breed, Age: age}
}
func (d Dog) GetAge() int {
	return d.Age
}
func (d *Dog) SetAge(age int) {
	d.Age = age
}
func (d Dog) GetInfo() string {
	return fmt.Sprintf("Name: %s, Breed: %s, Age: %d", d.Name, d.Breed, d.Age)
}

func main() {
	dog1 := NewDog("Buddy", "Golden Retriever", 3)
	fmt.Println(dog1.GetInfo())

	fmt.Println("Возраст:", dog1.GetAge())
	dog1.SetAge(5)
	fmt.Println("Новый возраст:", dog1.GetAge())
	fmt.Println(dog1.GetInfo())

	dog2 := NewDog("Bella", "Labrador", 7)
	fmt.Println(dog2.GetInfo())
}
