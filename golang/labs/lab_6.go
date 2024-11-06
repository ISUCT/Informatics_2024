package labs

import (
	"fmt"
)

type Rabbit struct {
	age int
}

func NewRabbit(age int) *Rabbit {
	r := new(Rabbit)
	r.age = age
	return r
}

func (r *Rabbit) getAge() int {
	return r.age
}

func (r *Rabbit) setAge(age int) int {
	r.age = age
	return r.age
}

func Lab6() {
	rabbit := NewRabbit(0)
	rabbit.setAge(17)
	age := rabbit.getAge()
	fmt.Printf("The age of the rabbit is %d years old", age)
}
