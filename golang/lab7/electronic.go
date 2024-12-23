package lab7

type Electronic struct {
	name  string
	price float64
	brand string
}

func (e *Electronic) GetName() string {
	return e.name
}

func (e *Electronic) GetPrice() float64 {
	return e.price
}

func (e *Electronic) SetPrice(price float64) {
	e.price = price
}

func (e *Electronic) ApplyDiscount(discount float64) {
	e.price -= e.price * discount / 100
}
