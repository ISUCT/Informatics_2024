package lab7

import "fmt"

type Vegetables struct {
	Name   string
	Weight float64
	Price  float64
}

func (c *Vegetables) make_sale(discount float64) {
	c.Price = c.Price * (1 - discount/100)
}
func (c *Vegetables) get_price() float64 {
	return c.Price
}
func (c *Vegetables) get_productInfo() string {
	return fmt.Sprintf("Name: %s, Weight %s, Price: %.2f", c.Name, c.Weight, c.Price)
}
