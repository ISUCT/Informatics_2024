package LABA7

type Dress struct {
	Name  string
	Price float64
	Brand string
	Color string
}

func (d *Dress) GetName() string {
	return d.Name
}

func (d *Dress) GetPrice() float64 {
	return d.Price
}

func (d *Dress) SetPrice(price float64) {
	d.Price = price
}

func (d *Dress) ApplyDiscount(discount float64) {
	d.Price = d.Price * (1 - discount/100)
}

func (d *Dress) ChangeColor(newColor string) {
	d.Color = newColor
}
