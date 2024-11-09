package labs

type Food struct {
	price    float64
	foodtype string
	brand    string
}

func (f *Food) getPrice() float64 {
	return f.price
}

func (f *Food) discount(discount float64) {
	f.price -= f.price * discount / 100
}

func (f *Food) setPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Food) changeCharacteristic(newFoodtype string, newBrand string) {
	f.foodtype = newFoodtype
	f.brand = newBrand
}
