package Lab7

type Bicycle struct {
	Name     string
	Brand    string
	Price    float64
	BikeType string
}

func (b *Bicycle) GetName() string {
	return b.Name
}

func (b *Bicycle) GetPrice() float64 {
	return b.Price
}

func (b *Bicycle) SetPrice(price float64) {
	b.Price = price
}

func (b *Bicycle) ApplyDiscount(discount float64) {
	b.Price -= b.Price * discount / 100
}

func (b *Bicycle) ChangeBikeType(newBikeType string) {
	b.BikeType = newBikeType
}
