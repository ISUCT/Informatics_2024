package lab6

import "fmt"

type Computer struct {
	Name  string
	Color string
	Capacity float64
}

func NewComputer(name, color string, Capacity float64) *Computer {
	c := new(Computer)
	c.Name = name
	c.Color = color
	c.Capacity = Capacity
	return c
}
func (c *Computer) SetCapacity(Capacity float64) { c.Capacity = Capacity }
func (c Computer) GetCapacity() float64       { return c.Capacity }
func (c Computer) GetColor() float64       { return c.Capacity }

func Runlab6() {
	DELL := NewComputer("ДЕЛЛ", "чёрный", 128.0)
	DELL.SetCapacity(256.0)
	fmt.Println(DELL.GetCapacity())
	fmt.Println(DELL.GetColor())
}