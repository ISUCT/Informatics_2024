package Lab7

type Shirt struct {
	Name     string
	Price    float64
	Brand    string
	Color    string
	Material string
	Size     string
}

func (sh *Shirt) GetName() string {
	return sh.Name
}

func (sh *Shirt) GetPrice() float64 {
	return sh.Price
}

func (sh *Shirt) GetBrand() string {
	return sh.Brand
}

func (sh *Shirt) SetPrice(price float64) {
	sh.Price = price
}

func (sh *Shirt) ApplyDiscount(discount float64) {
	sh.Price -= sh.Price * discount / 100
}

func (sh *Shirt) ChangeColor(newColor string) {
	sh.Color = newColor
}
func (sh *Shirt) ChangeMaterial(newMaterial string) {
	sh.Material = newMaterial
}

func (sh *Shirt) ChangeSize(newSize string) {
	sh.Size = newSize
}
