package laba7

type Product interface {
	SetDiscount(float64)
	SetPrice(float64)
	GetPrice() float64
}
