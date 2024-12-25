package lab7

import "errors"

var ErrColor error = errors.New("invalid color")
var ErrMaxHorsepower error = errors.New("the wrong maxHorsepower value")

type Color string

const (
	Gold  Color = "золотой"
	Chrome Color = "хром"
	White Color = "белый"
	Black Color = "чёрный"
)

type bibika struct {
	product
	color    Color
	maxHorsepower float64
}

func (c *bibika) SetColor(color Color) error {
	for _, standardColor := range []Color{Gold, Chrome, White, Black} {
		if color == standardColor {
			c.color = color
			return nil
		}
	}
	return ErrColor
}

func (c *bibika) SetHorsepower(maxHorsepower float64) error {
	if maxHorsepower < 0 || maxHorsepower > 1000 {
		return ErrMaxHorsepower
	}
	c.maxHorsepower = maxHorsepower
	return nil
}

func NewBibika(id int, name string, price float64, color Color, maxHorsepower float64) *bibika {
	c := &bibika{
		product: newProduct(id, name, price),
	}
	c.SetColor(color)
	c.SetHorsepower(maxHorsepower)
	return c
}
