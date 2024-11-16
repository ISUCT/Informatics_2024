package lab7

type Product interface {
	getPrice() float64
	setPrice(newPrice float64)
	discount(discount float64)
	setCharacteristics(string, string)
}
