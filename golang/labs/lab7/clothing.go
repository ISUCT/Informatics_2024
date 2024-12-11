package lab7

import (
	"fmt"
)

type Clothing struct {
	Name  string
	Price float64
	Size  string
	Color string
}

func (c *Clothing) applyDiscount(discount float64) error {
	if err := validateDiscount(discount); err != nil {
		return err
	}
	c.Price = c.Price * (1 - discount/100)
	return nil
}

func (c *Clothing) getPrice() float64 {
	return c.Price
}

func (c *Clothing) getProductInfo() string {
	if c.Name == "" || c.Size == "" || c.Color == "" {
		return ("Не все поля заполнены!!!")
	}
	return fmt.Sprintf("Name: %s, Size: %s, Color: %s, Price: %.2f", c.Name, c.Size, c.Color, c.Price)
}
