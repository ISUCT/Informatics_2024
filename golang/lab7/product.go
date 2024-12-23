package lab7

type Product interface {
	GetName() string
	GetPrice() float64
	ApplyDiscount(discount float64)
}

