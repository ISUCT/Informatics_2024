package labs

type Product interface {
	Setdiscount(discount float64)
	SetPrice(newPrice float64)
	GetPrice() float64
	GetName() string
	GetInfo() string
}
