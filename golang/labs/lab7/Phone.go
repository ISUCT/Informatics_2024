package lab7

type Phone struct {
	Size  float64
	Model string
	Price float64
}

func (p *Phone) SetSize(Size float64) {
	p.Size = Size
}

func (p *Phone) GetSize() float64 {
	return p.Size
}
func (p *Phone) SetModel(Model string) {
	p.Model = Model
}

func (p *Phone) SetPrice(Price float64) {
	p.Price = Price
}

func (p *Phone) GetPrice() float64 {
	return p.Price
}

func (p *Phone) SetDiscount(percent float64) {
	p.Price -= (p.Price / 100) * percent
}
