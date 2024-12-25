package lab7

type Products interface {
	GetName() string
	SetPrice(float64)
	GetPrice() float64
	GetInfo()
	ApplyDiscount(float64)
}
