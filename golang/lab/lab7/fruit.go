package lab7

import "fmt"

type Fruit struct {
	Name      string
	Freshness string
	Price     float32
}

func (f Fruit) GetInfo() {
	fmt.Println("У нас есть в наличии", f.Name, "оно", f.Freshness, "и стоит", f.Price)
}

func (f Fruit) GetPrice() float32 {
	return f.Price
}

func (f *Fruit) MakeDiscount(x float32) {
	(*f).Price = (f.Price / 100) * (100 - x)
}

func (f *Fruit) ChangePrice(x float32) {
	(*f).Price = x
}

func (f *Fruit) ChangeDescription(x string) {
	(*f).Freshness = x
}
