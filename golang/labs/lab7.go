package labs

import (
	"fmt"
)

type Product interface {
	sale(discount float64)
	price() float64
	productFeature() string
}

type Electronics struct {
	Name  string
	Price float64
	Brand string
	Model string
}

func (e *Electronics) sale(discount float64) {
	e.Price = e.Price * (1 - discount/100)
}

func (e *Electronics) price() float64 {
	return e.Price
}

func (e *Electronics) productFeature() string {
	return fmt.Sprintf("Name: %s, Brand: %s, Model: %s, Price: %.2f", e.Name, e.Brand, e.Model, e.Price)
}

type Clothing struct {
	Name  string
	Price float64
	Size  string
	Color string
}

func (c *Clothing) sale(discount float64) {
	c.Price = c.Price * (1 - discount/100)
}

func (c *Clothing) price() float64 {
	return c.Price
}

func (c *Clothing) productFeature() string {
	return fmt.Sprintf("Name: %s, Size: %s, Color: %s, Price: %.2f", c.Name, c.Size, c.Color, c.Price)
}

type Food struct {
	Name   string
	Price  float64
	Weight float64
}

func (f *Food) sale(discount float64) {
	f.Price = f.Price * (1 - discount/100)
}

func (f *Food) price() float64 {
	return f.Price
}

func (f *Food) productFeature() string {
	return fmt.Sprintf("Name: %s, Weight: %.2f kg, Price: %.2f", f.Name, f.Weight, f.Price)
}

func GeneralPrice(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.price()
	}
	return totalCost
}

func RunLab7() {
	product1 := &Electronics{Name: "Телефон", Price: 20, Brand: "Iphone", Model: "5s"}
	product1.sale(25)
	product2 := &Clothing{Name: "Рубашка", Price: 14.32, Size: "L", Color: "white"}
	product2.sale(10)
	product3 := &Food{Name: "Яблоко", Price: 2.35, Weight: 1}
	product3.sale(15)

	products := []Product{product1, product2, product3}

	for _, product := range products {
		fmt.Println(product.productFeature())
	}

	total := GeneralPrice(products)
	fmt.Printf("Общая стоимость: %.2f\n", total)
}
