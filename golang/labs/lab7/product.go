package labs

type Product interface {
	SetDiscount(discount float64)
	SetPrice(newPrice float64)
	GetPrice() float64
	GetName() string
	GetInfo() string
}
