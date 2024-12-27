package lab7

import "fmt"

type Electronics struct {
	Name  string
	Firm  string
	Price float64
}

func (e Electronics) GetInfo() {
	fmt.Println("В наличии есть", e.Name, "от фирмы", e.Firm, "и цену", e.Price)
}

func (e Electronics) GetPrice() float64 {
	return e.Price
}

func (e *Electronics) ApplyDiscount(x float64) {
	(*e).Price = (e.Price / 100) * (100 - x)
}

func (e *Electronics) EditPrice(x float64) {
	(*e).Price = x
}

func (e *Electronics) EditDescription(x string) {
	(*e).Firm = x
}
