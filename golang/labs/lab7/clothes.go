package labs

import "fmt"

type Clothes struct {
	name     string
	price    float64
	brand    string
	material string
	season   string
}

func (c *Clothes) GetName() string {
	return c.name
}

func (c *Clothes) GetPrice() float64 {
	return c.price
}

func (c *Clothes) SetDiscount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Clothes) SetPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) GetBrand() string {
	return c.brand
}
func (c *Clothes) GetSeason() string {
	return c.season
}
func (c *Clothes) GetMaterial() string {
	return c.material
}
func (c *Clothes) GetInf() string {
	return fmt.Sprintf("название: %s, цена: %.2f, бренд: %s, материал: %s,сезон: %s", c.name, c.price, c.brand, c.material, c.season)
}
