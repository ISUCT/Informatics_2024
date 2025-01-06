package lab7

import "fmt"

type Microwave_oven struct {
	Name      string
	Price     float64
}

func (f Microwave_oven) GetDescription() {
	fmt.Println("Микроволновая печь", f.Name, "стоимостью", f.Price)
}

func (f Microwave_oven) GetPrice() float64 {
	return f.Price
}

func (f *Microwave_oven) ApplyDiscount(x float64) {
	(*f).Price = (f.Price) * (x)
}

func (f *Microwave_oven) SetNewPrice(x float64) {
	(*f).Price = x
}

func (f *Microwave_oven) SetNewDescription(x string) {
	(*f).Name = x
}
