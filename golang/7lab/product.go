package sevenlab

type Product interface {
	applyDiscount(discount float64) error
	getPrice() float64
	getProductInfo() string
}
