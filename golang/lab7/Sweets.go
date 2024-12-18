package lab7
import "fmt"
type Sweets struct {
	Name  string
	Taste  string
	Price float32
}

func (s Sweets) GetInformation() {
	fmt.Println("У нас есть в наличии", s.Name, "вкуса", s.Taste, "стоимостью", s.Price)
}

func (s Sweets) GetPrice() float32 {
	return s.Price
}

func (s *Sweets) ApplyDiscount(x float32) {
	(*s).Price = (s.Price / 100) * (100 - x)
}

func (s *Sweets) ChangePrice(x float32) {
	(*s).Price = x
}

func (s *Sweets) ChangeDescription(x string) {
	(*s).Taste = x
}
