package lab6

import (
	"fmt"
)

type cat struct {
	name string
	age int
	breed string
}

func (c *cat) updateAge(new_age int) {
	(*c).age = new_age
}

func (c *cat) updateStruct (new_name string, new_age int, new_breed string) {
	(*c).name = new_name
	(*c).age = new_age
	(*c).breed = new_breed
}

func (c cat) printAge() int {
	return c.age
}

func (c cat) printName() string {
	return c.name
}

func (c cat) printBreed() string {
	return c.breed
}

func RunLab6Tasks() {
	var Tom cat = cat{"Tom", 9, "Scottish Fold"}
	var TomPointer *cat = &Tom
	fmt.Println("У нас есть кот по имени", Tom.printName(), "возрастом", Tom.printAge(), "лет и породы", Tom.printBreed())
	TomPointer.updateAge(13)
	fmt.Println("Через 4 года коту", Tom.printName(), "будет", Tom.printAge(), "лет")
	TomPointer.updateStruct("Misa", 5, "Maine coon")
	fmt.Println("Также у нас есть кошка по имени", Tom.printName(), "возрастом", Tom.printAge(), "лет и породы", Tom.printBreed())
}