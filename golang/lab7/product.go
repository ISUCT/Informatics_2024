package lab7

type Product interface {
	GetName() string
	SetName(NewName string)
	GetPrice() float64
	SetPrice(NewPrice float64)
	ApplyDiscount(discount float64)
}
