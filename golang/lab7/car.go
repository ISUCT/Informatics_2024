package lab7

import "errors"

type car struct {
	product
	color    string
	maxSpeed float64
}

func (c *car) SetColor(color string) error {
	for _, standardColor := range []string{"красный", "белый", "зелённый", "синий"} {
		if color == standardColor {
			c.color = color
			return nil
		}
	}
	return errors.New("invalid color")
}

func (c *car) SetSpeed(maxSpeed float64) error {
	if maxSpeed < 0 || maxSpeed > 250 {
		return errors.New("the wrong maxSpeed value")
	}
	c.maxSpeed = maxSpeed
	return nil
}

func NewCar(id int, name string, price float64, color string, maxSpeed float64) *car {
	c := &car{
		product: newProduct(id, name, price),
	}
	c.SetColor(color)
	c.SetSpeed(maxSpeed)
	return c
}
