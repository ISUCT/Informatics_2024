package lab7

type Phone struct {
	Name    string
	Price   float64
	Brand   string
	Storage float64
}

func (p *Phone) GetName() string {
	return p.Name
}

func (p *Phone) GetPrice() float64 {
	return p.Price
}

func (p *Phone) GetBrand() string {
	return p.Brand
}

func (p *Phone) ApplyDiscount(discount float64) {
	p.Price -= p.Price * discount / 100
}

func (p *Phone) ChangeStorage(newStorageSize float64) {
	p.Storage = newStorageSize
}
