package lab6

import "fmt"

type Phone struct {
	Color  string
	OS     string
	Number string
}

func NewPhone(color, os, number string) *Phone {
	p := new(Phone)
	p.Color = color
	p.OS = os
	p.Number = number
	return p
}
func (p *Phone) SetNumber(number string) { p.Number = number }
func (p Phone) GetNumber() string        { return p.Number }
func (p Phone) GetOS() string            { return p.OS }

func Lab6() {
	samsung := NewPhone("Чёрный", "Андроид", "89105789216")
	samsung.SetNumber("89204035907")
	fmt.Println(samsung.GetNumber())
	fmt.Println(samsung.GetOS())
}
