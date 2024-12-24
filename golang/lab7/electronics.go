package lab7

type Electronics struct {
	name          string
	price         float64
	brand         string
	model         string
	warranty      int
	isRefurbished bool
}

func (e Electronics) GetName() string {
	return e.name
}

func (e Electronics) GetPrice() float64 {
	return e.price
}

func (e *Electronics) SetPrice(price float64) {
	e.price = price
}

func (e *Electronics) ApplyDiscount(discount float64) {
	e.price -= e.price * discount
}
