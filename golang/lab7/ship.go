package lab7

type Ship struct {
	name     string
	material string
	price    float64
}

func (s *Ship) GetName() string {
	return s.name
}

func (s *Ship) SetName(NewName string) {
	s.name = NewName
}

func (s *Ship) GetPrice() float64 {
	return s.price
}

func (s *Ship) SetPrice(NewPrice float64) {
	s.price = NewPrice
}

func (s *Ship) ApplyDiscount(discount float64) {
	s.price = s.price * (100 - discount) / 100
}

func (s *Ship) SetMaterial(NewMaterial string) {
	s.material = NewMaterial
}
