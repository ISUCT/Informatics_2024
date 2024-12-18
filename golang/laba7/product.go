package laba7

type Product interface {
	ApplyDiscount(discount float64)
	GetPrice() float64
	GetProductInfo() string
}
