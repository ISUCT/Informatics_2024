package lab7

type Sneakers struct {
	Name  string
	Price float64
	Brand string
	Color string
}

func (s *Sneakers) GetName() string {
	return s.Name
}

func (s *Sneakers) GetPrice() float64 {
	return s.Price
}

func (s *Sneakers) SetPrice(price float64) {
	s.Price = price
}

func (s *Sneakers) ApplyDiscount(discount float64) {
	s.Price -= s.Price * discount / 100
}
