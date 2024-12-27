package lab7

type Brush struct {
	name     string
	price    float64
	material string
}

func (b *Brush) GetName() string {
	return b.name
}
func (b *Brush) GetPrice() float64 {
	return b.price
}
func (b *Brush) SetPrice(price float64) {
	b.price = price
}
func (b *Brush) ApplyDiscount(discount float64) {
	b.price -= b.price * discount / 100
}
