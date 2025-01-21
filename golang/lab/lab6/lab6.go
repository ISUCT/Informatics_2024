package tests

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
func inf() {
	myDog := NewDog("Buddy", "Labrador", 3)
	fmt.Println(myDog.GetInfo())
	fmt.Println("Age:", myDog.GetAge())
	myDog.SetAge(4)
	fmt.Println(myDog.GetInfo())
}
