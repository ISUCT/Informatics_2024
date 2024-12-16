package labs

import "fmt"

type Food struct {
	price    float64
	name     string
	brand    string
	calories int
}

func (f *Food) GetName() string {
	return f.name
}
func (f *Food) GetPrice() float64 {
	return f.price
}

func (f *Food) Setdiscount(discount float64) {
	f.price -= f.price * discount / 100
}

func (f *Food) SetPrice(newPrice float64) {
	f.price = newPrice
}
func (f *Food) GetInfo() string {
	return fmt.Sprintf("Название: %s, Цена: %.2f, Калорийность: %d , Бренд: %s", f.name, f.price, f.calories, f.brand)
}
