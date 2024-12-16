package lab7

import "fmt"

type Fruit struct {
	Name      string
	Freshness string
	Price     float64
}

func (f Fruit) GetInfo() {
	fmt.Println("У нас есть в наличии", f.Name, "оно", f.Freshness, "и стоит", f.Price)
}

func (f Fruit) GetPrice() float64 {
	return f.Price
}

func (f *Fruit) MakeDiscount(x float64) {
	(*f).Price = (f.Price / 100) * (100 - x)
}

func (f *Fruit) ChangePrice(x float64) {
	(*f).Price = x
}

func (f *Fruit) ChangeDescription(x string) {
	(*f).Freshness = x
}
