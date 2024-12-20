package lab7

type Food struct {
	price    float64
	name     string
	calories float64
}

func (f *Food) getName() string {
	return f.name
}

func (f *Food) setName(newName string) {
	f.name = newName
}

func (f *Food) getPrice() float64 {
	return f.price
}

func (f *Food) setPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Food) setCalories(newCalories float64) {
	f.calories = newCalories
}

func (f *Food) applyDiscount(discount float64) {
	f.price = f.price * (100 - discount) / 100
}
