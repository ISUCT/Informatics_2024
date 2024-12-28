package lab7

import "fmt"

type Fastfood struct {
	price    float64
	name     string
	brand    string
	calories int
}

func (f *Fastfood) getName() string {
	return f.name
}
func (f *Fastfood) getPrice() float64 {
	return f.price
}

func (f *Fastfood) setdiscount(discount float64) {
	f.price -= f.price * discount / 100
}

func (f *Fastfood) setPrice(newPrice float64) {
	f.price = newPrice
}
func (f *Fastfood) GetInfo() string {
	return fmt.Sprintf("Продается товар: %s, ценой: %.2f, имеющий калорийность: %d , бренда: %s", f.name, f.price, f.calories, f.brand)
}
