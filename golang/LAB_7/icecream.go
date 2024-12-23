package LAB_7

type IceCream struct {
	Name   string
	Brand  string
	Price  float64
	Flavor string
}

func (i *IceCream) GetName() string {
	return i.Name
}

func (i *IceCream) GetPrice() float64 {
	return i.Price
}

func (i *IceCream) SetPrice(price float64) {
	i.Price = price
}

func (i *IceCream) ApplyDiscount(discount float64) {
	i.Price -= i.Price * discount / 100
}

func (i *IceCream) ChangeFlavor(newFlavor string) {
	i.Flavor = newFlavor
}
