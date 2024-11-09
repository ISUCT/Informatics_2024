package lab7

import "errors"

var ErrColor error = errors.New("invalid color")
var ErrMaxSpeed error = errors.New("the wrong maxSpeed value")

type Color string

const (
	Red   Color = "красный"
	Green Color = "зелённый"
	White Color = "белый"
	Black Color = "чёрный"
)

type car struct {
	product
	color    Color
	maxSpeed float64
}

func (c *car) SetColor(color Color) error {
	for _, standardColor := range []Color{Red, Green, White, Black} {
		if color == standardColor {
			c.color = color
			return nil
		}
	}
	return ErrColor
}

func (c *car) SetSpeed(maxSpeed float64) error {
	if maxSpeed < 0 || maxSpeed > 250 {
		return ErrMaxSpeed
	}
	c.maxSpeed = maxSpeed
	return nil
}

func NewCar(id int, name string, price float64, color Color, maxSpeed float64) *car {
	c := &car{
		product: newProduct(id, name, price),
	}
	c.SetColor(color)
	c.SetSpeed(maxSpeed)
	return c
}
