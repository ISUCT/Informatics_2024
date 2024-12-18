package laba7

type Product interface {
	ProductInfo() string
	ApplyDiscount(discount float64)
	GetPrice() float64
}
