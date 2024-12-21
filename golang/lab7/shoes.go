package lab7

import "fmt"

type Shoes struct {
	Name  string
	Size  string
	Price float64
}

func (s *Shoes) GetInfo() {
	fmt.Println("Есть в наличии", s.Name, "размера", s.Size, "стоимостью", s.Price)
}

func (s *Shoes) GetPrice() float64 {
	return s.Price
}

func (s *Shoes) MakeDiscount(discountPercentage float64) {
    if discountPercentage >= 100 {
        s.Price = 0
        return
    }
    s.Price = s.Price * (1 - discountPercentage/100)
}

func (s *Shoes) ChangePrice(newPrice float64) {
	s.Price = newPrice
}

func (s *Shoes) ChangeDescription(newSize string) {
	s.Size = newSize
}
