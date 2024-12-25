package lab7

import "fmt"

type Food struct {
	Name   string
	Weight float64
	Price  float64
}

func (f *Food) make_sale(discount float64) {
	f.Price = f.Price * (1 - discount/100)
}
func (f *Food) get_price() float64 {
	return f.Price
}
func (f *Food) get_productInfo() string {
	return fmt.Sprintf("Name: %s, Weight: %.2f, Price: %.2f", f.Name, f.Weight, f.Price)
}
