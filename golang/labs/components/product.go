package components

type Product interface {
	sale(discount float64) error
	price() float64
	productFeature() string
}
