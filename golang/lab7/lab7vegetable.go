package lab7

type vegetable struct {
	Name  string
	Price float64
	Sort string
}

func (v *vegetable) GetName() string {
	return v.Name
}

func (v *vegetable) GetPrice() float64 {
	return v.Price
}

func (v *vegetable) SetPrice(price float64) {
	v.Price = price
}

func (v *vegetable) ApplyDiscount(discount float64) {
	v.Price -= v.Price * discount / 100
}
func (v *vegetable) GetSort(newSort string) {
	v.Sort = newSort
}