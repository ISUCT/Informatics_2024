package lab7

import "fmt"

type Clothes struct {
	Name  string
	Price float64
	Size  string
	Color string
}

func (c *Clothes) sale(discount float64) {
	c.Price = c.Price * (1 - discount/100)
}
func (c *Clothes) price() float64 {
	return c.Price
}
func (c *Clothes) productInfo() string {
	return fmt.Sprintf("Name: %s, Size: %s, Color: %s, Price: %.2f", c.Name, c.Size, c.Color, c.Price)
}

type Food struct {
	Name   string
	Weight float64
	Price  float64
}

func (f *Food) sale(discount float64) {
	f.Price = f.Price * (1 - discount/100)
}
func (f *Food) price() float64 {
	return f.Price
}
func (f *Food) productInfo() string {
	return fmt.Sprintf("Name: %s, Weight: %.2f, Price: %.2f", f.Name, f.Weight, f.Price)
}

type Vehicle struct {
	Brand  string
	Model  string
	Color  string
	MphMax float64
	Price  float64
}

func (v *Vehicle) sale(discount float64) {
	v.Price = v.Price * (1 - discount/100)
}
func (v *Vehicle) price() float64 {
	return v.Price
}
func (v *Vehicle) productInfo() string {
	return fmt.Sprintf("Brand: %s, Model: %s, Color: %s, MphMax: %.2f, Price: %.2f", v.Brand, v.Model, v.Color, v.MphMax, v.Price)
}

func priceAllProducts(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.price()
	}
	return totalCost
}