package labs

type Food struct {
	price    float64
	name     string
	brand    string
	calories int
}

func (f *Food) getName() string {
	return f.name
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
func (f *Food) getBrand() string {
	return f.brand
}

func (f *Food) getCalories() int {
	return f.calories
}
