package Lab7

import (
	"fmt"
)

type Cloth struct {
	Name  string
	Price float64
	Size  string
	Color string
}

func (c *Cloth) applyDiscount(discount float64) error {
	if err := validateDiscount(discount); err != nil {
		return err
	}
	c.Price = c.Price * (1 - discount/100)
	return nil
}

func (c *Cloth) getPrice() float64 {
	return c.Price
}

func (c *Cloth) getProductInfo() string {
	if c.Name == "" || c.Size == "" || c.Color == "" {
		return ("Не все поля заполнены")
	}
	return fmt.Sprintf("Название: %s, Размер: %s, Цвет: %s, Цена: %.2f", c.Name, c.Size, c.Color, c.Price)
}
