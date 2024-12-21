package lab7

type Fruits struct {
	price    float64
	name     string
	calories float64
}

func (f *Fruits) GetName() string {
	return f.name
}

func (f *Fruits) SetName(newName string) {
	f.name = newName
}

func (f *Fruits) GetPrice() float64 {
	return f.price
}

func (f *Fruits) SetPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Fruits) SetCalories(newCalories float64) {
	f.calories = newCalories
}

func (f *Fruits) ApplyDiscount(discount float64) {
	f.price = f.price * (100 - discount) / 100
}
