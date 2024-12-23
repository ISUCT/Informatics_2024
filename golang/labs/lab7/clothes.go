package lab7

import "fmt"

type Clothes struct {
	name  string
	price float64
	color string
	size  string
}

func NewClothes(name string, price float64, color string, size string) *Clothes {
	if price <= 0 {
		price = 1
	}
	c := &Clothes{name: name, price: price, color: color, size: size}
	return c
}

func (c *Clothes) applyDiscount(discount float64) error {
	err := validateDiscount(discount)
	if err != nil {
		return err
	}
	c.price = c.price * (1 - (discount / 100))
	return nil
}

func (c *Clothes) setPrice(newPrice float64) error {
	err := validatePrice(newPrice, c.price)
	if err != nil {
		return err
	}
	c.price = newPrice
	return nil
}

func (c *Clothes) getPrice() float64 {
	return c.price
}

func (c *Clothes) ChangeClothesSize(newSize string) error {
	if newSize == c.size {
		return fmt.Errorf("размер не поменялся")
	}
	c.size = newSize
	return nil
}

func (c *Clothes) printInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, color: %s\n, size: %s", c.name, c.price, c.color, c.size)
}
