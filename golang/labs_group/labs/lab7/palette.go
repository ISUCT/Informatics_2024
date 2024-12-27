package lab7

type Palette struct {
	name       string
	price      float64
	colorCount int
}

func (p *Palette) GetName() string {
	return p.name
}

func (p *Palette) GetPrice() float64 {
	return p.price
}

func (p *Palette) SetPrice(price float64) {
	p.price = price
}

func (p *Palette) ApplyDiscount(discount float64) {
	p.price -= p.price * discount / 100
}
