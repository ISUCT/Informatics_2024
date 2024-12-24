package lab7

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	GetInfo()
	ApplyDiscount(discount float64)
}
