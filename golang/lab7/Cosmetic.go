package lab7
import "fmt"

type Cosmetic struct {
	Name   string
	Brend string
	Price  float32
}

func (c Cosmetic) GetInformation() {
	fmt.Println("У нас есть в наличии", c.Name, "из бренда", c.Brend, "и стоит", c.Price)
}

func (c Cosmetic) GetPrice() float32 {
	return c.Price
}

func (c *Cosmetic) ApplyDiscount(x float32) {
	(*c).Price = (c.Price / 100) * (100 - x)
}

func (c *Cosmetic) ChangePrice(x float32) {
	(*c).Price = x
}

func (c *Cosmetic) ChangeDescription(x string) {
	(*c).Brend = x
}
