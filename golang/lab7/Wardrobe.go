package LABA7

type Wardrobe struct {
	Name  string
	Price float64
	Brand string
	Color string
}

func (w *Wardrobe) GetName() string {
	return w.Name
}

func (w *Wardrobe) GetPrice() float64 {
	return w.Price
}

func (w *Wardrobe) SetPrice(price float64) {
	w.Price = price
}

func (w *Wardrobe) ApplyDiscount(discount float64) {
	w.Price = w.Price * (1 - discount/100)
}

func (w *Wardrobe) ChangeColor(newColor string) {
	w.Color = newColor
}
