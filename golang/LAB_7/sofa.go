package LAB_7

type Sofa struct {
	Name     string
	Price    float64
	Brand    string
	Color    string
	Material string
}

func (s *Sofa) GetName() string {
	return s.Name
}

func (s *Sofa) GetPrice() float64 {
	return s.Price
}

func (s *Sofa) SetPrice(price float64) {
	s.Price = price
}

func (s *Sofa) ApplyDiscount(discount float64) {
	s.Price -= s.Price * discount / 100
}

func (s *Sofa) ChangeColor(newColor string) {
	s.Color = newColor
}
func (s *Sofa) ChangeMaterial(newMaterial string) {
	s.Material = newMaterial
}
