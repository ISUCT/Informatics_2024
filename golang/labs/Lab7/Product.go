package lab7


type Product interface {
	sale(discount float64)
	price() float64
	productInfo() string
}