package lab6

import (
	"fmt"
)

type Plane struct {
	Name   string
	Color  string
	Speed  float64
	Height float64
}

func NewPlane(name, color string, speed, height float64) *Plane {
	c := new(Plane)
	c.Name = name
	c.Color = color
	c.Speed = speed
	c.Height = height
	return c
}
func (c *Plane) SetSpeed(speed float64)   { c.Speed = speed }
func (c Plane) GetSpeed() float64         { return c.Speed }
func (c *Plane) SetHeigth(height float64) { c.Height = height }
func (c Plane) GetHeigth() float64        { return c.Height }
func (c Plane) GetColor() string          { return c.Color }
func CompleteLaba6() {
	boing := NewPlane("Боинг", "Белый", 800.0, 10000.0)
	boing.SetSpeed(900.0)
	boing.SetHeigth(12000.0)
	fmt.Println(boing.GetSpeed())
	fmt.Println(boing.GetHeigth())
	fmt.Println(boing.GetColor())
}
