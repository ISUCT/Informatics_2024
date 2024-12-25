package lab7

import "fmt"

type Vehicle struct {
	Brand string
	Model string
	Color string
	Price float64
}

func (v *Vehicle) make_sale(discount float64) {
	v.Price = v.Price * (1 - discount/100)
}
func (v *Vehicle) get_price() float64 {
	return v.Price
}
func (v *Vehicle) get_productInfo() string {
	return fmt.Sprintf("Brand: %s, Model: %s, Color: %s, Price: %.2f", v.Brand, v.Model, v.Color, v.Price)
}
