package lab7

type Mobile struct {
	Name  string
	Price float64
	Brand string
	Color string
}

func (m *Mobile) GetName() string {
	return m.Name
}

func (m *Mobile) GetPrice() float64 {
	return m.Price
}

func (m *Mobile) SetPrice(price float64) {
	m.Price = price
}

func (m *Mobile) ApplyDiscount(discount float64) {
	m.Price -= m.Price * discount / 100
}