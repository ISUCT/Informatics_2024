package lab7

import "fmt"

type Vegetables struct {
	Name  string
	Fresh string
	Price float64
}

func (v Vegetables) GetInfo() {
	fmt.Println("В наличии есть", v.Name, "который", v.Fresh, "и цену", v.Price)
}

func (v Vegetables) GetPrice() float64 {
	return v.Price
}

func (v *Vegetables) ApplyDiscount(x float64) {
	(*v).Price = (v.Price / 100) * (100 - x)
}

func (v *Vegetables) EditPrice(x float64) {
	(*v).Price = x
}

func (v *Vegetables) EditDescription(x string) {
	(*v).Fresh = x
}
