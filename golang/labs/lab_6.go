package labs

import (
	"fmt"
)  

type Rabbit struct {  
	Name   string 
	Age    int
	Weight float64
}  

func NewRabbit(name string, age int, weight float64) Rabbit {
  return Rabbit{Name: name, Age: age, Weight: weight}
}

func (r *Rabbit) GetAge() int {
  return r.Age
}

func (r *Rabbit) SetAge(age int) {
	if (age > 0) {
		r.Age = age
	}
}

func (r Rabbit) Info() string {
  return fmt.Sprintf("Имя: %s, Возраст: %d, Вес: %.2f кг", r.Name, r.Age, r.Weight)
}

func RunLab6() {
  rabbit := NewRabbit("Артём", 2, 1.5)

  fmt.Println(rabbit.Info())

  rabbit.SetAge(3)

  fmt.Printf("Обновленный возраст: %d\n", rabbit.GetAge())

  fmt.Println(rabbit.Info())
}