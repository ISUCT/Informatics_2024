package components

import (
	"errors"
	"fmt"
)

type Clothing struct {
	Name  string
	Price float64
	Size  string
	Color string
}

func (c *Clothing) sale(discount float64) error {
	if c.Name == "" || c.Size == "" || c.Color == "" {
		return errors.New("Не все поля заполнены для Clothing")
	}
	if discount < 0 {
		return errors.New("Скидка не может быть меньше 0")
	}
	if discount > 100 {
		return errors.New("Скидка не может больше 100%")
	}
	c.Price = c.Price * (1 - discount/100)
	return nil
}

func (c *Clothing) price() float64 {
	return c.Price
}

func (c *Clothing) productFeature() string {
	if c.Name == "" || c.Size == "" || c.Color == "" {
		return ("Не все поля заполнены!!!")
	}
	return fmt.Sprintf("Name: %s, Size: %s, Color: %s, Price: %.2f", c.Name, c.Size, c.Color, c.Price)
}
