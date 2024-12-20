package lab7

type Electronic struct {
	name        string
	price       float64
	description string
}

func (e *Electronic) GetPrice() float64 {
	return e.price
}

func (e *Electronic) GetName() string {
	return e.name
}

func (e *Electronic) ApplyDiscount(discount float64) {
	e.price -= e.price * discount
}

func (e *Electronic) SetPrice(price float64) {
	e.price = price
}

func (e *Electronic) SetName(name string) {
	e.name = name
}

func (e *Electronic) GetDescription() string {
	return e.description
}

func (e *Electronic) SetDescription(description string) {
	e.description = description
}
