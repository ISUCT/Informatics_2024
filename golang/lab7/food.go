package lab7

type Food struct {
	price       float64
	section     string
	nameproduct string
}

func (f *Food) getPrice() float64 {
	return f.price
}

func (f *Food) setPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Food) discount(discount float64) {
	f.price = f.price * (100 - discount) / 100
}

func (f *Food) setCharacteristics(newSection string, newNameproduct string) {
	f.section = newSection
	f.nameproduct = newNameproduct
}
