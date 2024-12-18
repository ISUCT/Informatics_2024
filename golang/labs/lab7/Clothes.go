package lab7

import "fmt"

type Clothes struct {
	Name  string
	Price float64
	Size  string
	Color string
}

func (c *Clothes) make_sale(discount float64) {
	c.Price = c.Price * (1 - discount/100)
}
func (c *Clothes) get_price() float64 {
	return c.Price
}
func (c *Clothes) get_productInfo() string {
	return fmt.Sprintf("Name: %s, Size: %s, Color: %s, Price: %.2f", c.Name, c.Size, c.Color, c.Price)
}
