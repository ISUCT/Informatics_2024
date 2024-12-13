package labs

import (
	"fmt"
)

type Rabbit struct {
	age    int
	name   string
	status string
}

func NewRabbit(age int, name string, status string) *Rabbit {
	r := new(Rabbit)
	r.age = age
	r.name = name
	r.status = status
	return r
}

func (r *Rabbit) getInfo() {
	fmt.Println("The name of the rabbit is", r.name)
	fmt.Printf("The age of the rabbit is %d years old\n", r.age)
	fmt.Println("The rabbit are", r.status)
}

func (r *Rabbit) changeStatus(status string) string {
	r.status = status
	fmt.Println("The rabbit are", r.status, "now")
	return r.status
}

func (r *Rabbit) grow() int {
	r.age = (r.age + 1)
	fmt.Println("The age of the rabbit is", r.age, "years old now")
	return r.age
}

func Lab6() {
	rabbit := NewRabbit(17, "Johnny", "sleep")
	rabbit.getInfo()
	fmt.Printf("\n")
	rabbit.changeStatus("jump")
	rabbit.grow()
	fmt.Printf("\n")
	rabbit.getInfo()
}
