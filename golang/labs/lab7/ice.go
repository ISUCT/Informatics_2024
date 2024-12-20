package lab7

import "fmt"

type Icecream struct {
	Name  string
	Price float64
	Taste string
}

func NewIcecream(name string, price float64, taste string) *Icecream {
	icecream := &Icecream{Name: name, Price: price, Taste: taste}
	return icecream
}

func (i *Icecream) GetName() string {
	return i.Name
}

func (i *Icecream) GetPrice() float64 {
	return i.Price
}

func (i *Icecream) SetPrice(price float64) {
	i.Price = price
}

func (i *Icecream) GetTaste() string {
	return i.Taste
}

func (i *Icecream) ApplyDiscount(discount float64) {
	i.Price -= i.Price * discount / 100
}

func (i *Icecream) GetInfo() {
	fmt.Printf("name: %v\nprice: %.2f\ntaste: %v\n", i.Name, i.Price, i.Taste)
}
