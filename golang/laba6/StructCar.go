package laba6

type Car struct {
	Name  string
	Color string
	Speed float64
}

func NewCar(name, color string, speed float64) *Car {
	c := new(Car)
	c.Name = name
	c.Color = color
	c.Speed = speed
	return c
}
func (c *Car) SetSpeed(speed float64) { c.Speed = speed }
func (c Car) GetSpeed() float64       { return c.Speed }
func (c Car) GetColor() string        { return c.Color }
