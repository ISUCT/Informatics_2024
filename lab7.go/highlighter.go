package lab7

type highlighter struct {
	price float32
	color string
	material string
}

func (x highlighter) GetPrice() float32 {
	return x.price
}

func (c *highlighter) Discount(x float32) {
 (*c).price = (c.price / 100) * (100 - x)
}

func (c *highlighter) ChangePrice(x float32) {
 (*c).price = x
}

func (c *highlighter) ChangeCharacteristic(newcolor string,newmaterial string ) {
	c.color = newcolor
	c.material = newmaterial
}

