package lab7

import "fmt"

type Book struct {
	Name   string
	Price  float64
	Format string
}

func (c *Book) make_sale(discount float64) {
	c.Price = c.Price * (1 - discount/100)
}
func (c *Book) get_price() float64 {
	return c.Price
}
func (c *Book) get_productInfo() string {
	return fmt.Sprintf("Name: %s, Format %s, Price: %.2f", c.Name, c.Format, c.Price)
}
