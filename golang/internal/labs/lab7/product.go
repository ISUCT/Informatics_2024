package lab7

type Product interface {
	GetPrice() float64
	GetName() string
	ApplyDiscount(discount float64)
	SetPrice(price float64)
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
}
