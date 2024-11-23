package lab7

type Furniture struct {
	price    float64
	name     string
	material string
}

func (f *Furniture) getPrice() float64 {
	return f.price
}

func (f *Furniture) setPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Furniture) applyDiscount(discount float64) {
	f.price = f.price * (100 - discount) / 100
}

func (f *Furniture) getName() string {
	return f.name
}

func (f *Furniture) setName(newName string) {
	f.name = newName
}

func (f *Furniture) setMaterial(newMaterial string) {
	f.material = newMaterial
}
