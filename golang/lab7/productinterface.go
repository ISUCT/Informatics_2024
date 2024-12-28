package lab7

type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	SetDiscount(discount float64)
}
