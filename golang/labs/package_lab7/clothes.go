package labs

type Clothes struct {
	name     string
	price    float64
	brand    string
	material string
	season   string
}

func (c *Clothes) getName() string {
	return c.name
}
func (c *Clothes) getPrice() float64 {
	return c.price
}

func (c *Clothes) discount(discount float64) {
	c.price -= c.price * discount / 100
}

func (c *Clothes) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) getBrand() string {
	return c.brand
}

func (c *Clothes) getSeason() string {
	return c.season
}
func (c *Clothes) getMaterial() string {
	return c.material
}
