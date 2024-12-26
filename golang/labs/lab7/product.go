package lab7

type Product interface {
	setdiscount(discount float64)
	setPrice(newPrice float64)
	getPrice() float64
	getName() string
	GetInfo() string
}
