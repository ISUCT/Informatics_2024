package laba7

type Product interface {
	applyDiscount(discount float64)
	getPrice() float64
	getProductInfo() string
}
