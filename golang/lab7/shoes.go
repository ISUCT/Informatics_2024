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

func (s *Shoes) MakeDiscount(x float64) {
    if x >= 100 {
        s.Price = 0
        return
    }
    s.Price = s.Price * (1 - x/100)
}

func (s *Shoes) ChangePrice(x float64) {
	s.Price = x
}

func (s *Shoes) ChangeDescription(x string) {
	s.Size = x
}