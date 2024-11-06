package lab6

import "fmt"

type Dog struct {
  Age int
}

func NewAge(age int) *Dog {
  x := new(Dog)
  x.Age = age
  return x
}

func (p *Dog) SetAge(age int) { p.Age = age }

func (p Dog) GetName() int   { return p.Age}

func RunLab6() {
  dog := NewAge(2)
  dog.SetAge(7)
  fmt.Println(dog.GetName())
}