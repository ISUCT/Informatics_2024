package lab7

type Fruit struct {
	name   string
	price  float64
	weight float64
}

func (f *Fruit) GetName() string {
	return f.name
}

func (f *Fruit) GetPrice() float64 {
	return f.price
}

func (f *Fruit) SetPrice(price float64) {
	f.price = price
}

func (f *Fruit) ApplyDiscount(discount float64) {
	f.price -= f.price * discount / 100
}
