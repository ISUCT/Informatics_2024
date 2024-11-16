package lab7

type Furniture struct {
	price         float64
	tupefurniture string
	material      string
}

func (f *Furniture) getPrice() float64 {
	return f.price
}

func (f *Furniture) setPrice(newPrice float64) {
	f.price = newPrice
}

func (f *Furniture) discount(discount float64) {
	f.price = f.price * (100 - discount) / 100
}

func (f *Furniture) setCharacteristics(newTupefurniture string, newMaterial string) {
	f.tupefurniture = newTupefurniture
	f.material = newMaterial
}
